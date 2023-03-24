package webx

import (
	"fmt"
	"net"
	"net/http"
)

// 评论：采用动词 Handle 更符合 Go 的命名风格
type HandleFunc func(ctx *Context)

type Server interface {
	http.Handler

	// Start 启动服务器
	// addr 是监听地址，如果只指定端口，可以使用 “:8081”，或者 “localhost:8081”
	Start(addr string) error 

	// AddRoute 注册一个路由
	// method 是 HTTP 方法
	// path 是路径，必须以 / 开头
	AddRoute(method, path string, handler HandleFunc)
}


// 编译，确保一定实现了 Server 接口
var _ Server = &HTTPServer{}


type HTTPServer struct {
	router
}

// HTTPServerImpl 也许更直观点
func NewHTTPServer() *HTTPServer {
	return &HTTPServer{
		router: newRouter(),
	}
}

// 评论：
// 1. 言简意赅，不像 gin 核心接口 IRoutes 中的 Handle 模棱两可，看上去像是处理什么东西，而实质上只是注册路由；
// 2. 此处还省去了 Get、Post 等方法，包裹一层又何必呢？简洁
func (h *HTTPServer) AddRoute(method, path string, handler HandleFunc) {
	h.router.addRoute(method, path, handler)
}


// ServeHTTP HTTPServer 处理请求的入口
func (h *HTTPServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := &Context{
		Req:  r,
		Resp: w,
	}

	// 接下来，查找路由、执行业务逻辑
	h.serve(ctx)
}


func (h *HTTPServer) Start(addr string) error {
	// 端口启动前
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	// 端口启动后
	fmt.Println("成功监听端口：", addr)

	/*
		这块儿做服务发现 - 注册本服务器到管理平台，譬如 etcd
	*/

	return http.Serve(listener, h)

	// 评论：监听端口和服务器启动分离，原本阻塞，不方便塞些处理操作，故毙掉。 
	// return http.ListenAndServe(addr, h)
}


func (h *HTTPServer) serve(ctx *Context) {
	n, ok := h.findRoute(ctx.Req.Method, ctx.Req.URL.Path)
	fmt.Println(ctx.Req.Method, ctx.Req.URL.Path)

	// 没找到 或者有路径，没注册业务逻辑，有屌用
	if !ok || n.handler == nil {
		ctx.Resp.WriteHeader(http.StatusNotFound)
		ctx.Resp.Write([]byte("404 了"))
		return
	}

	n.handler(ctx)
}


