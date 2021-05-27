package server

import (
	"commons"
	"commons/manager"
)

func Register() {
	registerOnLineServer()
}

// registerOnLineServer 注册网络服务
func registerOnLineServer() {
	manager.NetWorkRegister(func() commons.NetWork {
		return newNetServer()
	})
}
