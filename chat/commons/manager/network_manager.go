package manager

import (
	"commons"
	"commons/config"
	"context"
	"github.com/panjf2000/gnet"
	log "github.com/sirupsen/logrus"
)

var netWorks func() commons.NetWork

type NetworkManager struct {
	transform commons.Transforms

	con commons.NetWork //连接

	list chan *[]byte

	ctx    context.Context
	cancel context.CancelFunc
}

func (obj *NetworkManager) Setup() error {
	obj.list = make(chan *[]byte, 1024)
	obj.ctx, obj.cancel = context.WithCancel(context.Background())

	tmp := netWorks()
	if err := tmp.Setup(config.NetworkConfig.Network, config.NetworkConfig.IP, config.NetworkConfig.Port, func(data *[]byte, con gnet.Conn) error {
		obj.list <- data
		return nil
	}); err != nil {
		return err
	}
	go tmp.Run()
	obj.con = tmp
	return nil
}

func (obj *NetworkManager) SetSink(transform commons.Transforms) {
	obj.transform = transform
}

func (obj *NetworkManager) Run() {
	// 处理信息接口 可以开n个协程处理n个连接
The:
	for {
		select {
		case <-obj.ctx.Done():
			break The
		case data := <-obj.list:
			if err := obj.transform.Line(data); err != nil {
				log.Errorf("UDPManager Run line fail,err=%v \n", err)
			}
		}
	}
	//通知退出
	obj.transform.Stop()
}

func (obj *NetworkManager) Stop() {
	obj.cancel()
}

func (obj *NetworkManager) UnInit() {
	obj.con.Stop()
}
