package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	userClient "mall_api/rpc/user/client"
)

func MyOrder(ctx *gin.Context) {
	client := userClient.NewUserApiRpcClientWithContext(ctx)
	reply, err := client.Client.MyOrder(ctx, &userClient.MyOrderReq{
		UserId: cast.ToInt64(ctx.Param("id")),
	})
	if err != nil {
		ctx.JSON(0, err.Error())
		return
	}
	ctx.JSONP(200, reply)
}

func Info(ctx *gin.Context) {
	client := userClient.NewUserApiRpcClientWithContext(ctx)
	reply, err := client.Client.Info(ctx, &userClient.InfoReq{
		UserId: cast.ToInt64(ctx.Param("id")),
	})
	if err != nil {
		ctx.JSON(0, err.Error())
		return
	}
	ctx.JSONP(200, reply)
}
