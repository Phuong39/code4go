package manager

import (
	"chat_server/common"
	"chat_server/common/config"
	"chat_server/common/crontab"
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cast"
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
	create func() common.EventHandler //创建函数
	sortId int                        //创建序号
}

var creater []createStruct

type EventManager struct {
	sink      common.Sink
	chans     []chan *common.EventLine
	eventChan []chan Event
	handler   [][]common.EventHandler
	group     sync.WaitGroup
	ctx       context.Context
	cancel    context.CancelFunc
}

func (obj *EventManager) SetSink(sink common.Sink) {
	obj.sink = sink
}

func (obj *EventManager) Index(line *common.EventLine) int {
	data := fmt.Sprintf("%v-%v", line.AppID, line.PCID)
	idx := crc32.ChecksumIEEE([]byte(data))
	return int(idx % uint32(len(obj.chans))) //防止uint32 >> int出现负值
}

func (obj *EventManager) PushLine(line *common.EventLine) error {
	idx := obj.Index(line)
	obj.chans[idx] <- line
	return nil
}

func (obj *EventManager) Setup() error {
	obj.ctx, obj.cancel = context.WithCancel(context.Background())
	var count = config.BusinessConfig.EventCount
	obj.chans = make([]chan *common.EventLine, count)
	obj.eventChan = make([]chan Event, count)
	for i := 0; i < int(count); i++ {
		obj.chans[i] = make(chan *common.EventLine, 32)
		obj.eventChan[i] = make(chan Event, 32)
	}

	//排序
	sort.Slice(creater, func(i, j int) bool { return creater[i].sortId < creater[j].sortId })
	for count > 0 {
		var tmp []common.EventHandler
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

	spec1 := config.BusinessConfig.OnlineDataClearTime.Second + " " + config.BusinessConfig.OnlineDataClearTime.Minute +
		" " + config.BusinessConfig.OnlineDataClearTime.Hour + " * * *"
	err := crontab.RegisterCrontab(spec1, func() {
		for index := range obj.handler {
			obj.eventChan[index] <- clearJob
		}
	})
	if err != nil {
		log.Errorf("crontab RegisterCrontab fail,err=%v", err)
		return err
	}

	// */150 * * * * ?
	spec2 := "*/" + cast.ToString(config.BusinessConfig.KeepAlivePeriod) + " * * * * *"
	err = crontab.RegisterCrontab(spec2, func() {
		for index := range obj.handler {
			obj.eventChan[index] <- timeOut
		}
	})
	if err != nil {
		log.Errorf("crontab RegisterCrontab fail,err=%v", err)
		return err
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
							common.EventLinePool.Put(line)
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
