module mall_api

go 1.16

replace utils => ../utils

require (
	github.com/gin-gonic/gin v1.7.2
	github.com/golang/protobuf v1.4.3
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	github.com/opentracing/opentracing-go v1.2.0
	github.com/spf13/cast v1.4.0
	google.golang.org/grpc v1.39.0
	utils v0.0.0-00010101000000-000000000000
)
