
```go
/* 以下是示范代码用例 */

package main

import (
    webx "github.com/sjxiang/go-wheel/webx/v1"
)

func main() {
    
    h := webx.NewHTTPServer()

	h.AddRoute(http.MethodGet, "/user", func(ctx *webx.Context) {
		fmt.Println("处理第一件事")
		fmt.Println("处理第二件事")
	})

	handler1 := func(ctx *webx.Context) {
		fmt.Println("处理第一件事")
	}

	handler2 := func(ctx *webx.Context) {
		fmt.Println("处理第二件事")
	}

	h.AddRoute(http.MethodGet, "/user", func(ctx *webx.Context) {
		handler1(ctx)
		handler2(ctx)
	})

	h.Start(":8081")
}
```