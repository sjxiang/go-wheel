

### v2 - 前缀路由树


```go

/* 以下是示范代码用例 */

package main

import (
    webx "github.com/sjxiang/go-wheel/webx/v2"
)

func main() {
    
    h := webx.NewHTTPServer()

	h.AddRoute(http.MethodGet, "/home", func(ctx *webx.Context) {
		ctx.Resp.Write([]byte("hello, world"))
	})

	user := h.Group("/user")
	{
		user.AddRoute(http.MethodGet, "login", func(ctx *webx.Context) {
			ctx.Resp.Write([]byte("hello, world"))
		})

	}

	h.Start(":8081")

}

```