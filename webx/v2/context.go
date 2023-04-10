package webx

import "net/http"


// 代表上下文
type Context struct {
	Req    *http.Request
	Resp   http.ResponseWriter
	Params map[string]string
}