package main

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/sjxiang/go-wheel/webx"
)


func main() {
	s := webx.NewHTTPServer()

	s.AddRoute(http.MethodGet, "/index", func(ctx *webx.Context) {
		ctx.Resp.Write([]byte("Hello Gopher."))
	})
	s.AddRoute(http.MethodGet, "/user/*", func(ctx *webx.Context) {
		uuid, _ := uuid.NewUUID()
		content := fmt.Sprintf("通配符 %v.", uuid)
		ctx.Resp.Write([]byte(content))
	})
	s.AddRoute(http.MethodGet, "/user/:id", func(ctx *webx.Context) {
		ctx.Resp.Write([]byte("参数路径"))
	})
	
	s.Start(":8081")
}