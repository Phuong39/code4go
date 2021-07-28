module sing

go 1.12

require (
	github.com/gin-gonic/gin v1.4.0
	github.com/opentracing/opentracing-go v1.2.0
	github.com/uber/jaeger-client-go v2.29.1+incompatible
	utils v0.0.0-00010101000000-000000000000
)

replace utils => ../utils
