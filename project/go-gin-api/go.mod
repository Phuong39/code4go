module go-gin-api

go 1.12

require (
	github.com/gin-contrib/pprof v1.2.1
	github.com/gin-gonic/gin v1.4.0
	github.com/golang/protobuf v1.4.3
	github.com/grpc-ecosystem/go-grpc-middleware v1.1.0
	github.com/opentracing/opentracing-go v1.2.0
	github.com/xinliangnote/go-util v0.0.0-20191115235314-e7d1576d41db
	golang.org/x/time v0.0.0-20190921001708-c4c64cad1fd0
	google.golang.org/grpc v1.39.0
	gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc // indirect
	gopkg.in/gomail.v2 v2.0.0-20160411212932-81ebce5c23df // indirect
	utils v0.0.0-00010101000000-000000000000
)

replace utils => ../utils
