package handler

import "context"
import proto "yidu-user/protos"

type UserHandler struct{}

func (u *UserHandler) Login(ctx context.Context, req *proto.LoginReq, rsp *proto.LoginRsp) error {
	rsp.Greeting = "Hello " + req.Name
	return nil
}
