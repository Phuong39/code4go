package manager

import (
	"commons"
	"context"
	"errors"
	log "github.com/sirupsen/logrus"
)

var protocolSink []commons.TransformLine

type ProtocolManager struct {
	sink   commons.Sink
	ctx    context.Context
	cancel context.CancelFunc
}

func (obj *ProtocolManager) SetSink(sink commons.Sink) {
	obj.sink = sink
}

func (obj *ProtocolManager) Line(data *[]byte) error {
	var err error
	var eventLine *commons.EventLine
	if eventLine, err = obj.Transform(data); err != nil {
		log.Errorf("Transform line fail,err=%v \n", err)
	} else if err = obj.sink.PushLine(eventLine); err != nil {
		log.Errorf("PushLine to sink fail,err=%v \n", err)
	}
	return err
}
func (obj *ProtocolManager) Setup() error {
	obj.ctx, obj.cancel = context.WithCancel(context.Background())
	return nil
}

func (obj *ProtocolManager) Run() {
	// 通知退出
	// obj.sink.Stop()
}

// Transform 根据协议版本转换EventLine
func (obj *ProtocolManager) Transform(netWorkData *[]byte) (line *commons.EventLine, err error) {
	for _, t := range protocolSink {
		if line, err = t.Line(netWorkData); err == nil {
			return
		}
	}
	return nil, errors.New("unknown protocol")
}
func (obj *ProtocolManager) Stop() {
	obj.cancel()
	//阻塞: 外层会等待Run退出 ，所以这里可以没有wait都可以
}

func (obj *ProtocolManager) UnInit() {

}
