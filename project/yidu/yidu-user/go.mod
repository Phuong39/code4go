module yidu-user

go 1.16

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/golang/protobuf v1.5.2
	github.com/micro/go-micro/v2 v2.9.1
	google.golang.org/protobuf v1.26.0
	commons v0.0.0-00010101000000-000000000000
)

replace (
 	commons => ../commons
)
