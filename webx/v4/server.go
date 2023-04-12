package webx

import (
	"log"
	"net"
	"net/http"
)

// 注释：采用动词 Handle 更符合 Go 的命名风格
type HandleFunc func(ctx *Context)

// 通过编译检查，来确保一定实现了 Server 接口
var _ Server = &HTTPServer{}

type Server interface {
	http.Handler
	
	// Start 启动服务器
	// addr 是监听地址，如果只指定端口，可以使用 ":8081" 或者 "localhost:8081"
	Start(addr string) error

	// AddRoute 注册一个路由
	// method 是 HTTP 方法
	// path 是路由，必须以 / 开头
	// handleFunc 是你的业务逻辑
	AddRoute(method string, path string, handleFunc HandleFunc)
}

// 注释：ServerImpl，感觉更合适
type HTTPServer struct {
	// addr string 创建的时候传递，而不是 Start 接收。这个都是可以的

	router
	ms        []Middleware
}

func NewHTTPServer() *HTTPServer {
	return &HTTPServer{
		router: newRouter(),
	}
}

// 注册中间件
func (h *HTTPServer) Use(ms ...Middleware) {
	h.ms = ms
}

// ServeHTTP 处理请求的入口
func (h *HTTPServer) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	// 你的框架代码就在这里
	ctx := &Context{
		Req:  request,
		Resp: writer,
	}

	// 倒序构造
	root := h.serve
	for i := len(h.ms)-1; i >= 0; i-- {
		root = h.ms[i](root)

	}
	root(ctx)
}

func (h *HTTPServer) serve(ctx *Context) {
	// 接下来就是查找路由，并且执行命中的业务逻辑
	mi, ok := h.findRoute(ctx.Req.Method, ctx.Req.URL.Path)

	// 没找到，404
	if !ok || mi.n.handler == nil {
		ctx.Resp.WriteHeader(http.StatusNotFound)
		ctx.Resp.Write([]byte("Not Found"))
		return
	}

	ctx.Params = mi.segmentParams
	mi.n.handler(ctx)
}

// 评论：言简意赅，新增路由，而非 gin 中 IRoutes 的 Handle 表达的模棱两可；
// 看上去是在处理什么东西，而实际上只是注册路由
func (h *HTTPServer) AddRoute(method string, path string, handleFunc HandleFunc) {
	// 这里注册到路由树里面
	h.addRoute(method, path, handleFunc)
}


func (h *HTTPServer) Start(addr string) error {

	// 监听端口和启动 Server 分离

	l, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	log.Println("成功监听端口：", addr)
	
	/* 
	
		在这里，可以让用户注册所谓的 after start 回调
		比如说往你的 admin 注册一下自己这个实例
		在这里执行一些你业务所需的前置条件
	
		*/ 

	return http.Serve(l, h)
}
