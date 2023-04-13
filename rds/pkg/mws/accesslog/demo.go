package accesslog



// 存放系统中间件

// package middlewares

// import (
// 	"bytes"
// 	"fmt"
// 	"io/ioutil"
// 	"time"

// 	"github.com/gin-gonic/gin"
// 	"github.com/sjxiang/gohub/pkg/logger"
// 	"github.com/sjxiang/gohub/pkg/util"
// 	"go.uber.org/zap"
// )

// // 有些数据没啥接口访问，故包裹一层，拷贝一份
// type responseBodyWriter struct {
// 	gin.ResponseWriter
// 	body *bytes.Buffer
// }


// func (r responseBodyWriter) Write(b []byte) (int, error) {
// 	r.body.Write(b)
// 	return r.ResponseWriter.Write(b)
// }


// // Logger 记录请求日志
// func Logger() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
		
// 		// 获取 response 内容
// 		w := &responseBodyWriter{
// 			body: &bytes.Buffer{}, 
// 			ResponseWriter: ctx.Writer,
// 		}
// 		ctx.Writer = w
		
// 		// 获取请求数据
// 		var requestBody []byte
// 		if ctx.Request.Body != nil {

// 			// c.Request.Body 是一个 Buffer 对象，只能读取 1 次
// 			requestBody, _ := ioutil.ReadAll(ctx.Request.Body)

// 			// 读取后，重新赋值 c.Request.Body，以供后续的其他操作
// 			ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(requestBody))
		
// 		}

// 		// 设置开始时间
// 		start := time.Now()

// 		ctx.Next()

// 		// 开始记录日志的逻辑
// 		cost := time.Since(start)
// 		respStatus := ctx.Writer.Status()

// 		logFields := []zap.Field{
// 			zap.Int("status",  respStatus),
// 			zap.String("request", ctx.Request.Method+" "+ctx.Request.URL.String()),
// 			zap.String("query", ctx.Request.URL.RawQuery),
// 			zap.String("ip", ctx.ClientIP()),
// 			zap.String("user-agent", ctx.Request.UserAgent()),
// 			zap.String("errors", ctx.Errors.ByType(gin.ErrorTypePrivate).String()),
// 			zap.String("time", util.MicrosecondsStr(cost)),
// 		}

// 		if ctx.Request.Method == "POST" || ctx.Request.Method == "PUT" || ctx.Request.Method == "DELETE" {
			
// 			// 请求的内容
// 			logFields = append(logFields, zap.String("Request Body", string(requestBody)))

// 			// 响应的内容
// 			logFields = append(logFields, zap.String("Response Body", w.body.String()))

// 		}

		

// 		if respStatus > 400 && respStatus <= 499 {
// 			// 除了 400 StatusBadRequest 以外，warn 提示一下，常见的 404 开发时都要注意
// 			logger.Warn(
// 				fmt.Sprintf("HTTP 警告 %v",respStatus), 
// 				logFields...,
// 			)
// 		} else if respStatus >= 500 && respStatus <= 599 {
// 			// 除了内部错误，记录 error
// 			logger.Error(
// 				fmt.Sprintf("HTTP 错误 %v",respStatus), 
// 				logFields...,
// 			)
// 		} else {
// 			logger.Debug("HTTP 访问日志", logFields...)
// 		}
// 	}
// }
