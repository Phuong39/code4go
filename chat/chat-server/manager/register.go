package manager

import "chat_server/common"

// ProtocolRegister 协议转换器注册
func ProtocolRegister(transform common.TransformLine) error {
	protocolSink = append(protocolSink, transform)
	return nil
}

// EventRegister 事件处理器注册
// note 排序ID相同时，调用顺序不可预测
func EventRegister(sortID int, create func() common.EventHandler) error {
	creater = append(creater, createStruct{
		create, sortID,
	})
	return nil
}

// SinkRegister 输出注册（所有输出不允许再次修改line)
func SinkRegister(sink common.Sink) error {
	sinks = append(sinks, sink)
	return nil
}

// NetWorkRegister 输出network register
func NetWorkRegister(create func() common.NetWork) {
	netWorks = append(netWorks, create)
}
