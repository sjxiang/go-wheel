

### v3 - Context 处理输入输出


```go

/* 以下是示范代码用例 */
package main

import (
	"fmt"
	"net/http"

	webx "github.com/sjxiang/go-wheel/webx/v3"
)


func main() {
	s := webx.NewHTTPServer()

	s.AddRoute(http.MethodPost, "/user/register", func(ctx *webx.Context) {
		u := &User{}
		err := ctx.BindJSON(u)
		if err != nil {
			fmt.Println(err)
		}
		ctx.RespJSON(200, u.Name)
	})
	
	s.Start(":8081")
}


type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// $ curl -X POST  -H "Content-Type: application/json" -d '{"name": "Alice", "age": 30}'  http://localhost:8081/user/register

```