package webx

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strconv"
	"sync"
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

func (ctx *Context) QueryValueV1(key string) StringValue {
	if ctx.cacheQueryValue == nil {
		ctx.cacheQueryValue = ctx.Req.URL.Query()
	}

	vals, ok := ctx.cacheQueryValue[key]
	if !ok || len(vals) == 0 {
		return StringValue{err: errors.New("webx: key not found")}
	}

	return StringValue{val: vals[0]}
}


func (ctx *Context) QueryValueAsInt64V1(key string) (int64, error) {
	return ctx.QueryValueV1(key).ToInt64()
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


func (s StringValue) ToInt64() (int64, error) {
	if s.err != nil {
		return 0, s.err
	}

	return strconv.ParseInt(s.val, 10 , 64)
}




// JSON 响应
func (ctx *Context) RespJSON(code int, val any) error {
	bs, err := json.Marshal(val)
	if err != nil {
		return err
	}
	ctx.Resp.WriteHeader(code)
	_, err = ctx.Resp.Write(bs)
	return err
}


// 设置 cookie
func (ctx *Context) SetCookie(cookie *http.Cookie) {
	http.SetCookie(ctx.Resp, cookie)
}



// 线程安全
type SaveContext struct {
	c Context
	l sync.Mutex
}

// 装饰器
func (ctx *SaveContext) RespJSON(code int, val any) error {
	ctx.l.Lock()
	defer ctx.l.Unlock()

	return ctx.c.RespJSON(code, val)
}

