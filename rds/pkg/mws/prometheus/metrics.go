package prometheus

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

type MiddlewareBuilder struct {
	Namespace   string
	Subsystem   string
	ConstLabels map[string]string
	Name        string
	Help        string
}

func (m MiddlewareBuilder) Build() gin.HandlerFunc {
	
	// 采样
	summaryVec := prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name:        m.Name,
			Subsystem:   m.Subsystem,
			Namespace:   m.Name,
			Help:        m.Help,
			ConstLabels: m.ConstLabels,
			Objectives: map[float64]float64{
					0.5: 0.01,
					0.75: 0.01,
					0.90: 0.01,
					0.99: 0.001,
					0.999: 0.0001,
			},
		}, 
		[]string{"pattern", "method", "status"},
	)

	// 把观察者注册进去
	prometheus.MustRegister(summaryVec)

	return func(ctx *gin.Context) {
		startTime := time.Now()
		defer func() {
			duration := time.Now().Sub(startTime).Milliseconds()
			
			pattern := ctx.FullPath()
			if pattern == "" {
				pattern = "unknown"
			}
			
			statusCode := 200
			
			summaryVec.WithLabelValues(pattern, ctx.Request.Method, strconv.Itoa(statusCode)).Observe(float64(duration))
		}()

		ctx.Next()
	}
}


/*
	Prometheus

		响应码

		Vector 用法

		创建一个 Vector 向量，设置 ConstLabes 和 Labels

		使用 WithLabelValues 来设置具体的收集器
 */

