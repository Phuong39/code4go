module novel

go 1.15

require (
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/golang/protobuf v1.4.2
	github.com/jeanphorn/log4go v0.0.0-20190526082429-7dbb8deb9468
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-plugins/broker/nats v0.0.0-20200119172437-4fe21aa238fd
	github.com/micro/go-plugins/broker/nsq v0.0.0-20200119172437-4fe21aa238fd
	github.com/micro/go-plugins/broker/rabbitmq v0.0.0-20200119172437-4fe21aa238fd
	github.com/micro/go-plugins/broker/redis v0.0.0-20200119172437-4fe21aa238fd
	github.com/micro/go-plugins/registry/consul v0.0.0-20200119172437-4fe21aa238fd
	github.com/micro/go-plugins/registry/etcdv3 v0.0.0-20200119172437-4fe21aa238fd
	github.com/micro/go-plugins/transport/tcp v0.0.0-20200119172437-4fe21aa238fd
	github.com/toolkits/file v0.0.0-20160325033739-a5b3c5147e07 // indirect
	xorm.io/xorm v1.0.5
)
