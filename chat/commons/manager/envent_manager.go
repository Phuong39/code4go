package manager

import (
	"commons"
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"hash/crc32"
	"sort"
	"sync"
)

type Event uint8

const (
	timeOut = iota + 1
	clearJob
)

type createStruct struct {
	create func() commons.EventHandler //创建函数
	sortId int                         //创建序号
}

var creater []createStruct

var EventCount uint

type EventManager struct {
	sink      commons.Sink
	chans     []chan *commons.EventLine
	eventChan []chan Event
	handler   [][]commons.EventHandler
	group     sync.WaitGroup
	ctx       context.Context
	cancel    context.CancelFunc
}

func (obj *EventManager) SetSink(sink commons.Sink) {
	obj.sink = sink
}

func (obj *EventManager) Index(line *commons.EventLine) int {
	data := fmt.Sprintf("%v-%v", line.AppID, line.PCID)
	idx := crc32.ChecksumIEEE([]byte(data))
	return int(idx % uint32(len(obj.chans))) //防止uint32 >> int出现负值
}

func (obj *EventManager) PushLine(line *commons.EventLine) error {
	idx := obj.Index(line)
	obj.chans[idx] <- line
	return nil
}

func (obj *EventManager) Setup() error {
	obj.ctx, obj.cancel = context.WithCancel(context.Background())
	var count = EventCount
	obj.chans = make([]chan *commons.EventLine, count)
	obj.eventChan = make([]chan Event, count)
	for i := 0; i < int(count); i++ {
		obj.chans[i] = make(chan *commons.EventLine, 32)
		obj.eventChan[i] = make(chan Event, 32)
	}

	//排序
	sort.Slice(creater, func(i, j int) bool { return creater[i].sortId < creater[j].sortId })
	for count > 0 {
		var tmp []commons.EventHandler
		for _, f := range creater {
			tmp = append(tmp, f.create())
		}
		if len(tmp) > 0 {
			obj.handler = append(obj.handler, tmp)
		}
		count -= 1
	}

	if len(obj.chans) != len(obj.handler) {
		panic("EventHandler setup error")
	}
	return nil
}

func (obj *EventManager) Run() {
	goFunc := func(index int) {
		obj.group.Add(1)

	filter:
		for {
			select {
			case <-obj.ctx.Done():
				obj.group.Done()
				break filter
			case line := <-obj.chans[index]:
				for _, h := range obj.handler[index] {
					if h.Handler(line) {
						//各种处理(上线、版本变更)
						if !h.Filter(line) {
							//当返回false时，即被视为数据被过滤
							commons.EventLinePool.Put(line)
							continue filter
						}
					}
				}
				//下一节点推送
				if err := obj.sink.PushLine(line); err != nil {
					log.Errorf("sink PushLine fail,err=%v \n", err)
				}
			case e := <-obj.eventChan[index]:
				h := obj.handler[index][0]
				switch e {
				case timeOut:
					var offLineList = h.OffLineList()
					for _, v := range offLineList {
						//下一节点推送
						if err := obj.sink.PushLine(v); err != nil {
							log.Errorf("sink PushLine fail,err=%v \n", err)
						}
					}
				case clearJob:
					h.OnlineDataClear()
				}
			}
		}
	}
	for idx := range obj.chans {
		//很多协程
		go goFunc(idx) //避免协程内使用idx
	}
}

func (obj *EventManager) Stop() {
	obj.cancel()
	obj.group.Wait()
	obj.sink.Stop()
}
