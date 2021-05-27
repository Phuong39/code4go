package server

import (
	"chat_server/common"
	"chat_server/manager"
)

func init() {
	registerOnLineServer()
}

// registerOnLineServer 注册网络服务
func registerOnLineServer() {
	manager.NetWorkRegister(func() common.NetWork {
		return newNetServer()
	})
}
