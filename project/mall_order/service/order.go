package service

import (
	"context"
	"mall_order/client"
)

var storeList []*Order

type Order struct {
	ID     int64
	UserID int64
}

func init() {
	storeList = append(storeList, &Order{
		ID: 1, UserID: 1,
	}, &Order{
		ID: 2, UserID: 1,
	}, &Order{
		ID: 3, UserID: 1,
	}, &Order{
		ID: 4, UserID: 1,
	}, &Order{
		ID: 5, UserID: 2,
	}, &Order{
		ID: 6, UserID: 2,
	}, &Order{
		ID: 7, UserID: 2,
	}, &Order{
		ID: 8, UserID: 2,
	})
}

type OrderApiServerImpl struct {
}

func (h *OrderApiServerImpl) CompleteOrder(ctx context.Context, req *client.CompleteOrderReq) (*client.CompleteOrderReply, error) {
	return nil, nil
}

func (h *OrderApiServerImpl) GetUserOrder(ctx context.Context, req *client.GetUserOrderReq) (*client.GetUserOrderReply, error) {
	reply := client.GetUserOrderReply{}
	reply.Code = 200
	for _, v := range storeList {
		if v.UserID == req.UserId {
			reply.OrderId = append(reply.OrderId, v.ID)
		}
	}
	return &reply, nil
}
