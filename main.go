package main

import (
	"fmt"
	"net/http"
	"path"
	"mime/multipart"

	webx "github.com/sjxiang/go-wheel/webx/v4"
	"github.com/sjxiang/go-wheel/webx/v4/middleware/accesslog"
)


func main() {
	s := webx.NewHTTPServer()

	s.Use(accesslog.MiddlewareBuilder{}.Build())

	s.AddRoute(http.MethodPost, "/user/register", func(ctx *webx.Context) {
		u := &User{}
		err := ctx.BindJSON(u)
		if err != nil {
			fmt.Println(err)
		}
		ctx.RespJSON(200, u.Name)
	})

	s.AddRoute(http.MethodPost, "/upload", func(ctx *webx.Context) {
		f := webx.FileUploader{
			// <input type="file" name="file" />
			FileField: "file",
			DstPathFunc: func(fh *multipart.FileHeader) string {
				return path.Join("testdata", "upload", fh.Filename)
			},
		}.Handle()
	
		f(ctx)
	})
	
	
	s.Start(":8081")
}


type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// $ curl -X POST \     
//        -H "Content-Type: application/json" \
//        -d '{"name": "Alice", "age": 30}'  \
//        http://localhost:8081/user/register
