module jaegers

go 1.16

require (
	github.com/gin-gonic/gin v1.7.2
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/protobuf v1.4.3
	github.com/google/uuid v1.3.0 // indirect
	github.com/opentracing/opentracing-go v1.2.0
	github.com/pborman/uuid v1.2.1
	github.com/uber/jaeger-client-go v2.29.1+incompatible
	google.golang.org/grpc v1.39.0
	utils v0.0.0-00010101000000-000000000000
)

replace utils => ../utils
