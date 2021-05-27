package manager

import "commons"

// ProtocolRegister 协议转换器注册
func ProtocolRegister(transform commons.TransformLine) error {
	protocolSink = append(protocolSink, transform)
	return nil
}

// EventRegister 事件处理器注册
// note 排序ID相同时，调用顺序不可预测
func EventRegister(sortID int, create func() commons.EventHandler) error {
	creater = append(creater, createStruct{
		create, sortID,
	})
	return nil
}

// EventCountSet 业务数量设置
func EventCountSet(eventCount uint) {
	EventCount = eventCount
}

// SinkRegister 输出注册（所有输出不允许再次修改line)
func SinkRegister(sink commons.Sink) error {
	sinks = append(sinks, sink)
	return nil
}

// NetWorkRegister 输出network register
func NetWorkRegister(create func() commons.NetWork) {
	netWorks = create
}
