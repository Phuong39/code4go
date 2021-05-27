package manager

import (
	"chat_server/common"
	"chat_server/common/config"
	"context"
	"github.com/panjf2000/gnet"
	log "github.com/sirupsen/logrus"
)

var netWorks func() common.NetWork

type UDPManager struct {
	transform common.Transforms

	con common.NetWork //连接

	list chan *[]byte

	ctx    context.Context
	cancel context.CancelFunc
}

func (obj *UDPManager) Setup() error {
	obj.list = make(chan *[]byte, 1024)
	obj.ctx, obj.cancel = context.WithCancel(context.Background())

	tmp := netWorks()
	if err := tmp.Setup(config.ServerConfig.Network, config.ServerConfig.IP, config.ServerConfig.Port, func(data *[]byte, con gnet.Conn) error {
		obj.list <- data
		return nil
	}); err != nil {
		return err
	}
	go tmp.Run()
	obj.con = tmp
	return nil
}

func (obj *UDPManager) SetSink(transform common.Transforms) {
	obj.transform = transform
}

func (obj *UDPManager) Run() {
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

func (obj *UDPManager) Stop() {
	obj.cancel()
}

func (obj *UDPManager) Uninit() {
	obj.con.Stop()
}
