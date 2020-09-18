/**
* 作者：刘时明
* 时间：2019/10/20-16:14
* 作用：
 */
package handler

import (
	"context"
	pbDemo "novel/protocol/demo"
)

func (u *DemoImpl) SayHello(ctx context.Context, req *pbDemo.SayHelloReq, rsp *pbDemo.SayHelloRsp) error {
	rsp.Result.ErrCode = 0
	rsp.Result.ErrMsg = req.Name
	return nil
}
