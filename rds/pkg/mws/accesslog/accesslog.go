package accesslog

import "github.com/gin-gonic/gin"



func AccessLog() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// 全局中间件 or

		ctx.Next()

	}
}


