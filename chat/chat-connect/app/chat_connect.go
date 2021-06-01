package app

import (
	"chat_connect/app/handler"
	"chat_connect/app/protocol"
	"chat_connect/app/server"
	"commons/manager"
	"sync"
)

type ChatConnect struct {
	event    manager.EventManager
	protocol manager.ProtocolManager
	network  manager.NetworkManager

	//同步处理
	group sync.WaitGroup
}

func (obj *ChatConnect) Setup() (err error) {
	server.Register()
	protocol.Register()
	handler.Register()

	if err = obj.network.Setup(); err != nil {
		return
	}
	if err = obj.protocol.Setup(); err != nil {
		return
	}
	if err = obj.event.Setup(); err != nil {
		return
	}

	// 业务链串连
	obj.protocol.SetSink(&obj.event)
	obj.network.SetSink(&obj.protocol)
	return
}

func (obj *ChatConnect) Run() {
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
		obj.network.Run()
		obj.group.Done()
	}()
}

func (obj *ChatConnect) Stop() {
	obj.network.Stop()
	obj.group.Wait()
}

func (obj *ChatConnect) UnInit() {
	obj.network.Stop()
	obj.group.Wait()
}
