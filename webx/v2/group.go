package webx


// 辅助功能 - 路由分组
type Group struct {
	prefix string
	s      Server
}

// 非侵入式
func (g *Group) AddRoute(method string, path string, handleFunc HandleFunc) {
	g.s.AddRoute(method, g.prefix + path, handleFunc)
}


func (h *HTTPServer) Group(prefix string) *Group {
	return &Group{
		prefix: prefix,
	}
}

