module github.com/Terry-Mao/goim

go 1.15

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/Shopify/sarama v1.19.0 // indirect
	github.com/Shopify/toxiproxy v2.1.4+incompatible // indirect
	github.com/bilibili/discovery v1.0.1
	github.com/bsm/sarama-cluster v2.1.15+incompatible
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/eapache/go-resiliency v1.1.0 // indirect
	github.com/eapache/go-xerial-snappy v0.0.0-20180814174437-776d5712da21 // indirect
	github.com/eapache/queue v1.1.0 // indirect
	github.com/gin-gonic/gin v1.3.0
	github.com/gogo/protobuf v1.1.1
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/protobuf v1.4.2
	github.com/golang/snappy v0.0.0-20180518054509-2e65f85255db // indirect
	github.com/gomodule/redigo v2.0.0+incompatible
	github.com/google/uuid v1.0.0
	github.com/onsi/ginkgo v1.14.1 // indirect
	github.com/onsi/gomega v1.10.2 // indirect
	github.com/pierrec/lz4 v2.0.5+incompatible // indirect
	github.com/rcrowley/go-metrics v0.0.0-20181016184325-3113b8401b8a // indirect
	github.com/smartystreets/assertions v0.0.0-20180927180507-b2de0cb4f26d // indirect
	github.com/stretchr/testify v1.3.0
	github.com/ugorji/go/codec v0.0.0-20190204201341-e444a5086c43 // indirect
	github.com/zhenjl/cityhash v0.0.0-20131128155616-cdd6a94144ab
	golang.org/x/net v0.0.0-20200520004742-59133d7f0dd7
	google.golang.org/grpc v1.16.0
	gopkg.in/Shopify/sarama.v1 v1.19.0
)

replace (
	// cloud.google.com/go => github.com/googleapis/google-cloud-go v0.26.0
	// google.golang.org/appengine => github.com/golang/appengine v1.1.0
	google.golang.org/genproto => github.com/google/go-genproto v0.0.0-20180817151627-c66870c02cf8
	google.golang.org/grpc => github.com/grpc/grpc-go v1.16.0
)
