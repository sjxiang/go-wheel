package webx

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strconv"
)

// 代表上下文
type Context struct {
	Req    *http.Request
	Resp   http.ResponseWriter
	Params map[string]string

	// 缓存
	cacheQueryValue url.Values
}



func (ctx *Context) BindJSON(val any) error {
	if ctx.Req.Body == nil {
		return errors.New("webx: body 为 nil")
	}

	// json 解码填充到结构体指针
	decoder := json.NewDecoder(ctx.Req.Body)
	return decoder.Decode(val)
}


// 处理表单
func (ctx *Context) FormValue(key string) (string, error) {
	err := ctx.Req.ParseForm()
	if err != nil {
		return "", err
	}

	return ctx.Req.FormValue(key), nil 
}


func (ctx *Context) FormValueAsInt64(key string) (int64, error) {
	if err := ctx.Req.ParseForm(); err != nil {
		return 0, err
	}

	val := ctx.Req.FormValue(key)
	return strconv.ParseInt(val, 10, 64) 
}



// 查询参数
func (ctx *Context) QueryValue(key string) (string, error) {
	if ctx.cacheQueryValue == nil {
		ctx.cacheQueryValue = ctx.Req.URL.Query()
	}

	vals, ok := ctx.cacheQueryValue[key]
	if !ok || len(vals) == 0 {
		return "", errors.New("webx: 找不到这个 key")
	}

	return vals[0], nil 
}


func (ctx *Context) QueryValueAsInt64(key string) (int64, error) {
	val, err := ctx.QueryValue(key)
	if err != nil {
		return 0, err
	}

	return strconv.ParseInt(val, 10, 64)
}



// 路径参数
func (ctx *Context) pathValue(key string) (string, error) {
	val, ok := ctx.Params[key]
	if !ok {
		return "", errors.New("webx: 找不到这个 key")
	}

	return val, nil 
}



// 通用的处理方案（解决转换指定类型）
type StringValue struct {
	val string
	err error
}

func (s StringValue) String() (string, error) {
	return s.val, s.err
}

