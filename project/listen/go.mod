module listen

go 1.12

require (
	github.com/golang/protobuf v1.4.3
	golang.org/x/net v0.0.0-20200822124328-c89045814202
	google.golang.org/grpc v1.39.0
	utils v0.0.0-00010101000000-000000000000
)

replace utils => ../utils
