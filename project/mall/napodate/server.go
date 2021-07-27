package napodate

import (
	"context"
	"github.com/gin-gonic/gin"
	httptransport "github.com/go-kit/kit/transport/http"
	"net/http"
)

// NewHTTPServer 是一个很好的服务器
func NewHTTPServer(ctx context.Context, endpoints Endpoints) http.Handler {
	r := gin.Default()

	r.Use(commonMiddleware())

	disk := r.Group("/api")
	{
		v1 := disk.Group("/v1")
		{
			v1.GET("/status", func(ctx *gin.Context) {
				httptransport.NewServer(endpoints.StatusEndpoint,
					decodeStatusRequest, encodeResponse)
			})
			v1.GET("/get", func(ctx *gin.Context) {
				httptransport.NewServer(endpoints.StatusEndpoint,
					decodeGetRequest, encodeResponse)
			})
			v1.POST("/validate", func(ctx *gin.Context) {
				httptransport.NewServer(endpoints.StatusEndpoint,
					decodeValidateRequest, encodeResponse)
			})
		}
	}
	return r
}

func commonMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Add("Context-Type", "application/json")
	}
}
