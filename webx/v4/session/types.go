package session



/*

	问题

	1. 管理 Session：如创建、查找、销毁和刷新
	2. Session 存储和查找用户数据：考虑这些数据真实存储在什么地方
	3. Session id 存储和提取：和 HTTP 的请求响应打交道



store 需要支持

- 创建 session：例如用户登陆成功之后，为他创建一个 session。session 应该有过期时间
- 销毁 session：当用户退出登录的时候，要销毁掉数据。
- 查找：根据 HTTP 请求里面携带的 session id，验证 session id，并查找对应的 session 实例
- 刷新：如果用户持续保持活跃，那么 session 应该在这期间一直有效。


	*/


import (
	"context"
	"net/http"
)

// 存储和查找用户设置的数据
type Session interface {
	Get(ctx context.Context, key string) (string, error)
	// val 如果设置为类型 any，那么对应的 Redis 之类的实现，就需要考虑序列化的问题
	Set(ctx context.Context, key string, val string) error
	ID() string
}

// 管理 Session 本身
type Store interface {
	Generate(ctx *context.Context, id string) (Session, error)
	Remove(ctx *context.Context, id string) error
	Get(ctx *context.Context, id string) (Session, error)

	Refresh(ctx *context.Context, id string) error
}


type Propagator interface {
	// 将 session id 从 HTTP 请求中提取出来
	Extract(req *http.Request) (string, error)
	// 将 session id 注入到 HTTP 响应里面，幂等
	Inject(id string, writer http.ResponseWriter) error
	// 将 session id 从 HTTP 响应中删除（后端已经删了，没多大意义，可有可无）
	Remove(writer http.ResponseWriter) error
}