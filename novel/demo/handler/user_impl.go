/**
* 作者：刘时明
* 时间：2019/10/20-16:14
* 作用：
 */
package handler

import "github.com/micro/go-micro"

type DemoImpl struct {
	svc micro.Service
}

func NewDemoImp(svc micro.Service) *DemoImpl {
	return &DemoImpl{svc: svc}
}
