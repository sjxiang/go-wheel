package accesslog

import (
	"bytes"
	"encoding/json"
	"io/ioutil"

	"github.com/sjxiang/go-wheel/webx/v4"
)


type MiddlewareBuilder struct {
	logFunc func(accessLog []byte)
}


func (mb MiddlewareBuilder) Build() webx.Middleware {
	return func(next webx.HandleFunc) webx.HandleFunc {
		return func(ctx *webx.Context) {
			body, _ := ioutil.ReadAll(ctx.Req.Body)

			l := accessLog {
				HTTPMethod: ctx.Req.Method,
				Body: string(body),
			}
			// 可重复读
			ctx.Req.Body = ioutil.NopCloser(bytes.NewBuffer(body))
			bs, err := json.Marshal(l)
			if err == nil {
				mb.logFunc(bs)
			}
			next(ctx)
		}
	}
}
 

type accessLog struct {
	Host       string
	Route      string
	HTTPMethod string `json:"http_method"`
	Path       string
	Body       string
}


// func (b *MiddlewareBuilder) Build() webx.Middleware {

	// return func(next web.HandleFunc) web.HandleFunc {
	// 	return func(ctx *web.Context) {
	// 		defer func() {
	// 			l := accessLog{
	// 				Host:       ctx.Req.Host,
	// 				Route:      ctx.MatchedRoute,
	// 				Path:       ctx.Req.URL.Path,
	// 				HTTPMethod: ctx.Req.Method,
	// 			}
	// 			val, _ := json.Marshal(l)
	// 			b.logFunc(string(val))
	// 		}()
	// 		next(ctx)
	// 	}
	// }
// }