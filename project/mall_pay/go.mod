module mall_pay

go 1.16

require (
	github.com/golang/protobuf v1.5.2
	google.golang.org/grpc v1.39.0
	utils v0.0.0-00010101000000-000000000000
)

replace utils => ../utils
