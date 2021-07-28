package service

import (
	"context"
	"mall_pay/client"
)

type PayApiServerImpl struct {
}

func (h *PayApiServerImpl) PayStatus(ctx context.Context, req *client.PayStatusReq) (*client.PayStatusReply, error) {
	return nil, nil
}
