package opentelemetry

import (
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

/*

	tracing 链路追踪

	概念：
		tracer 表示记录 trace 的实例，一般来说，tracer 会有一个对应的接口

		span 代表 trace 中的一段，因此 trace 本身也可以看作是一个 span。
			span 本身是一个层级概念，因此有父子关系。
			一个 trace 的 span 可以看做是多叉树

	工具：
		开源实现 SkyWalking、Zipkin、Jeager

	支持：（摆脱第三方依赖）
		OpenTelemetry 与 Zipkin 和 Jeager 的结合，核心在于构造出一个 TracerProvider，并且调用 otel.SetTraceProvider


*/


const instrumentationName = "extra/opentelemetry"

type MiddlewareBuilder struct {
	Tracer trace.Tracer
}


func (m MiddlewareBuilder) Build() gin.HandlerFunc {
	if m.Tracer == nil {
		m.Tracer = otel.GetTracerProvider().Tracer(instrumentationName)
	}
	return func(ctx *gin.Context) {
		
		spanCtx, span := m.Tracer.Start(ctx.Request.Context(), "unknown")
		// span.End 执行之后，就意味着 span 本身已经确定无疑了，将不再变化
		defer span.End()

		span.SetAttributes(attribute.String("http.method", ctx.Request.Method))
		span.SetAttributes(attribute.String("http.path", ctx.Request.URL.Path[:256]))
		span.SetAttributes(attribute.String("http.host", ctx.Request.Host))


		// spanCtx 传递下去（方便串起来）
		ctx.Request = ctx.Request.WithContext(spanCtx)
		ctx.Next()	
	
		span.SetAttributes(attribute.Int("http.status", 1))  // 需要包裹缓存一份
	}
}

