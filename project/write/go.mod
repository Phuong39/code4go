module write

go 1.12

require (
	github.com/golang/protobuf v1.4.3
	github.com/opentracing/opentracing-go v1.2.0
	golang.org/x/net v0.0.0-20200822124328-c89045814202
	google.golang.org/grpc v1.39.0
	utils v0.0.0-00010101000000-000000000000
)

replace utils => ../utils
