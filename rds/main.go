package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"

	"go.opentelemetry.io/otel/exporters/zipkin"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.10.0"

	"go.opentelemetry.io/otel/exporters/jaeger"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"

	"github.com/sjxiang/go-wheel/rds/pkg/mws/opentelemetry"
	"github.com/sjxiang/go-wheel/rds/pkg/mws/prometheus"
)



func Init() {
	ctx := context.Background()

	// 构造 tracer 实例
	tracer := otel.GetTracerProvider().Tracer("extra/opentelemetry")

	// 如果传入的 ctx 已经和一个 span 绑定了，那么新的 span 就是老的 span 的儿子
	ctx, span := tracer.Start(ctx, "opentelemetry-demo", trace.WithAttributes(attribute.String("version", "1")))
	defer span.End()

	// 重置名字
	span.SetName("otel-demo")
	span.SetAttributes(attribute.Int("status", 200))
	span.AddEvent("马老师，发生甚么事了？他偷袭，不讲武德")
}



func main() {

	r := gin.New()

	tracer := otel.GetTracerProvider().Tracer("extra/opentelemetry")
	
	store := cookie.NewStore([]byte("secret"))

	r.Use(
		opentelemetry.MiddlewareBuilder{Tracer: tracer}.Build(),
		prometheus.MiddlewareBuilder{Namespace: "app", Subsystem: "user", Name: "http_response"}.Build(),
		sessions.Sessions("custom-session", store),  // 一进入 context，就塞进去
	)

	r.GET("/hello", func(ctx *gin.Context) {
		session := sessions.Default(ctx)  // 从 context 里面拿出来
		if session.Get("hello") != "world" {
			session.Set("hello", "world")
			session.Save()
		}

		ctx.JSON(http.StatusOK, gin.H{
			"hello": session.Get("hello"),
		})
	})

	r.GET("/test", func(ctx *gin.Context) {
		firstCtx, firstSpan := tracer.Start(ctx.Request.Context(), "First_layer")
		defer firstSpan.End()

		time.Sleep(1*time.Second)

		_, secondSpan := tracer.Start(firstCtx, "First_layer")
		defer secondSpan.End()

		time.Sleep(300 * time.Millisecond)
	})




	initZipkin()

	r.Run()
}



func initZipkin() {
	// 要注意这个端口，和 docker-compose 中的保持一致
	exporter, err := zipkin.New(
		"http://localhost:19411/api/v2/spans",
		zipkin.WithLogger(log.New(os.Stderr, "opentelemetry-demo", log.Ldate|log.Ltime|log.Llongfile)),
	)
	if err != nil {
		panic(err)
	}

	batcher := sdktrace.NewBatchSpanProcessor(exporter)
	tp := sdktrace.NewTracerProvider(
			sdktrace.WithSpanProcessor(batcher),
			sdktrace.WithResource(
				resource.NewWithAttributes(
					semconv.SchemaURL,
					semconv.ServiceNameKey.String("opentelemetry-demo"),
				),
			),
	)

	otel.SetTracerProvider(tp)
}


func initJeager() {
	url := "http://localhost:14268/api/traces"
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	if err != nil {
		panic(err)
	}
	tp := sdktrace.NewTracerProvider(
		// Always be sure to batch in production.
		sdktrace.WithBatcher(exp),
		// Record information about this application in a Resource.
		sdktrace.WithResource(
			resource.NewWithAttributes(
				semconv.SchemaURL,
				semconv.ServiceNameKey.String("opentelemetry-demo"),
				attribute.String("environment", "dev"),
				attribute.Int64("ID", 1),
			),
		),
	)

	otel.SetTracerProvider(tp)
}