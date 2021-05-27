package manager

import (
	"chat_server/common"
	"context"
	log "github.com/sirupsen/logrus"
)

var sinks []common.Sink

type SinkManager struct {
	line   chan *common.EventLine
	ctx    context.Context
	cancel context.CancelFunc
}

func (obj *SinkManager) PushLine(line *common.EventLine) error {
	obj.line <- line
	return nil
}

func (obj *SinkManager) Setup() (err error) {
	for _, sink := range sinks {
		if err = sink.Setup(); err != nil {
			return
		}
	}
	obj.line = make(chan *common.EventLine, 32)
	obj.ctx, obj.cancel = context.WithCancel(context.Background())
	return
}
func (obj *SinkManager) Run() {
	for {
		select {
		case <-obj.ctx.Done():
			return
		case line := <-obj.line:
			for _, s := range sinks {
				if err := s.PushLine(line); err != nil {
					log.Errorf("sink PushLine fail,err=%v \n", err)
				}
			}
		}
	}
}
func (obj *SinkManager) Stop() {
	obj.cancel()
}

func (obj *SinkManager) Uninit() {

}
