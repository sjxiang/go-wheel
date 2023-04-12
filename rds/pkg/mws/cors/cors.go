package cors

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


type MiddlewareBuilder struct {
	AllowOrigin string
}

// 跨域
func (m MiddlewareBuilder) Build() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		allowOrigin := m.AllowOrigin
		if allowOrigin == "" {
			allowOrigin = ctx.Request.Header.Get("Origin")
		}

		ctx.Writer.Header().Set("Access-Control-Allow-Origin", allowOrigin)  // * 放行所有
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		ctx.Writer.Header().Set("Content-Type", "application/json")
		
		if ctx.Request.Method == http.MethodOptions {
			ctx.AbortWithStatus(http.StatusNoContent)
			return 
		}

		ctx.Next()
	}
}