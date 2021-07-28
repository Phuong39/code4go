package service

import (
	"context"
	"mall_user/client"
	"time"
)

var storeMap map[int64]*User

type User struct {
	Name    string
	HeadImg string
}

func init() {
	storeMap = make(map[int64]*User)
	storeMap[1] = &User{
		Name:    "lsm",
		HeadImg: "123.png",
	}
	storeMap[2] = &User{
		Name:    "mast",
		HeadImg: "123.png",
	}
}

type UserApiServerImpl struct {
}

func (h *UserApiServerImpl) Info(ctx context.Context, req *client.InfoReq) (*client.InfoReply, error) {
	reply := client.InfoReply{}
	reply.Code = 200
	user, ok := storeMap[req.UserId]
	if ok {
		reply.Name = user.Name
		reply.HeadImage = user.HeadImg
	}
	time.Sleep(50 * time.Millisecond)
	return &reply, nil
}

func (h *UserApiServerImpl) MyOrder(ctx context.Context, req *client.MyOrderReq) (*client.MyOrderReply, error) {
	reply := client.MyOrderReply{}
	reply.Code = 200
	reply.Data = nil
	return &reply, nil
}
