package main

import (
	"fmt"
	"net/http"

	webx "github.com/sjxiang/go-wheel/webx/v3"
)


func main() {
	s := webx.NewHTTPServer()

	s.AddRoute(http.MethodGet, "/index", func(ctx *webx.Context) {
		ctx.Resp.Write([]byte("Hello Gopher."))
	})

	s.AddRoute(http.MethodGet, "/user/login", func(ctx *webx.Context) {
		ctx.Resp.Write([]byte("登录"))
	})

	s.AddRoute(http.MethodPost, "/user/register", func(ctx *webx.Context) {
		u := &User{}
		err := ctx.BindJSON(u)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(u.Name)
	})

	
	s.AddRoute(http.MethodGet, "/article/:id", func(ctx *webx.Context) {
		ctx.Resp.Write([]byte("Hello article" + ctx.Params["id"]))
	})
	
	s.Start(":8081")
}


type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}