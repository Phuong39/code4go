package client

import (
	"github.com/gin-gonic/gin"
	grpc_middeware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
	"sync"
	"utils"
)

type UserApiRpcClient struct {
	conn   *grpc.ClientConn
	Client UserApiClient
}

var (
	defaultClient *UserApiRpcClient
	once          = &sync.Once{}
)

//func NewUserApiRpcClient() *UserApiRpcClient {
//	once.Do(func() {
//		//	建立RPC连接
//		// ClientInterceptorWithParent
//		conn, err := grpc.Dial(utils.UserHost, grpc.WithInsecure(), grpc.WithUnaryInterceptor(utils.ClientInterceptor(utils.DefTracer, context.Background())))
//		if err != nil {
//			panic(err)
//		}
//		defaultClient = &UserApiRpcClient{conn: conn, Client: NewUserApiClient(conn)}
//	})
//	return defaultClient
//}

func NewUserApiRpcClientWithContext(ctx *gin.Context) *UserApiRpcClient {
	once.Do(func() {
		tracer, _ := ctx.Get("Tracer")
		parentSpanContext, _ := ctx.Get("ParentSpanContext")
		conn, err := grpc.DialContext(
			ctx,
			utils.UserHost,
			grpc.WithInsecure(),
			grpc.WithBlock(),
			grpc.WithUnaryInterceptor(grpc_middeware.ChainUnaryClient(utils.ClientInterceptorWithParent(tracer.(opentracing.Tracer),
				parentSpanContext.(opentracing.SpanContext)))),
		)
		if err != nil {
			panic(err)
		}
		defaultClient = &UserApiRpcClient{conn: conn, Client: NewUserApiClient(conn)}
	})
	return defaultClient
}
