package app

import (
	"chat_server/common/crontab"
	"chat_server/manager"
	"sync"
)

type ChatServer struct {
	event    manager.EventManager
	protocol manager.ProtocolManager
	sink     manager.SinkManager
	udp      manager.UDPManager

	//同步处理
	group sync.WaitGroup
}

func (obj *ChatServer) Setup() (err error) {
	//server.Register()
	//protocol.Register()
	//handler.Register()
	//sink.Register()

	if err = obj.udp.Setup(); err != nil {
		return
	}
	if err = obj.protocol.Setup(); err != nil {
		return
	}
	if err = obj.event.Setup(); err != nil {
		return
	}
	if err = obj.sink.Setup(); err != nil {
		return
	}

	// 业务链串连
	obj.event.SetSink(&obj.sink)
	obj.protocol.SetSink(&obj.event)
	obj.udp.SetSink(&obj.protocol)
	return
}

func (obj *ChatServer) Run() {
	go func() {
		obj.group.Add(1)
		obj.sink.Run()
		obj.group.Done()
	}()
	go func() {
		obj.group.Add(1)
		obj.event.Run()
		obj.group.Done()
	}()
	go func() {
		obj.group.Add(1)
		obj.protocol.Run()
		obj.group.Done()
	}()
	go func() {
		obj.group.Add(1)
		obj.udp.Run()
		obj.group.Done()
	}()
	go func() {
		obj.group.Add(1)
		crontab.CrontabRun()
		obj.group.Done()
	}()
}

func (obj *ChatServer) Stop() {
	obj.udp.Stop()
	obj.group.Wait()
}

func (obj *ChatServer) UnInit() {
	obj.udp.Stop()
	obj.group.Wait()
}
