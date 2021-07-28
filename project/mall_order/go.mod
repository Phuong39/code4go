module mall_order

go 1.16

replace utils => ../utils

require (
	github.com/golang/protobuf v1.4.3
	google.golang.org/grpc v1.39.0
	utils v0.0.0-00010101000000-000000000000
)
