/**
* 作者：刘时明
* 时间：2019/10/20-23:27
* 作用：
 */
package client

import (
	pbDemo "novel/protocol/demo"
	"novel/utils"
)

var (
	DemoClient pbDemo.DemoService
)

func Init() {
	DemoClient = pbDemo.NewDemoService("", utils.Client())
}
