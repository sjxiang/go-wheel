

### v3 - Context 处理输入输出


```go

/* 以下是示范代码用例 */

package main

import (
	"net/http"
	webx "github.com/sjxiang/go-wheel/webx/v3"
)

func main() {
	s := webx.NewHTTPServer()

	s.AddRoute(http.MethodGet, "/index", func(ctx *webx.Context) {
		ctx.Resp.Write([]byte("Hello Gopher."))
	})

	user := s.Group("/user")
	{
		user.AddRoute(http.MethodGet, "/login", func(ctx *webx.Context) {
			ctx.Resp.Write([]byte("登录"))
		})
		user.AddRoute(http.MethodGet, "/register", func(ctx *webx.Context) {
			ctx.Resp.Write([]byte("注册"))
		})
	}
	
	s.AddRoute(http.MethodGet, "/article/:id", func(ctx *webx.Context) {
		ctx.Resp.Write([]byte("Hello article" + ctx.Params["id"]))
	})
	
	s.Start(":8081")
}

```