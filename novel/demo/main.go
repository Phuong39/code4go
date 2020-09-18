/**
* 作者：刘时明
* 时间：2019/10/30-23:51
* 作用：
 */
package main

import (
	"novel/demo/conf"
	"novel/demo/handler"
	puDemo "novel/protocol/demo"
	"novel/utils"
)

func main() {
	// 初始化配置文件
	conf.InitConfig()

	config := []utils.Option{
		func(o *utils.Options) {
			o.RegisterType = utils.MDNSRegister
		},
	}
	svc, err := utils.InitService(conf.SvcName, config...)
	if err != nil {
		panic(err)
	}
	// 注册handler
	_ = puDemo.RegisterDemoServiceHandler(svc.Server(), handler.NewDemoImp(svc))

	// 启动服务
	if err = utils.Start(); err != nil {
		panic(err)
	}
	utils.Stop()
}
