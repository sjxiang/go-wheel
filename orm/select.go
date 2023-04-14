package orm

import (
	"context"
	"reflect"
	"strings"
)



// 用于构造 SELECT 语句
type Selector[T any] struct {
	tbl string 
}

// 实现链式调用
// From 指定表名，如果是空字符串，那么将会使用默认表名
func (s *Selector[T]) From(tbl string) *Selector[T] {
	s.tbl = tbl
	return s
}


func (s *Selector[T]) Where(predicate ...Predicate) *Selector[T] {

	panic("implement me")
}

func (s *Selector[T]) Get(ctx *context.Context) (*T, error) {
	panic("implement me")
}


func (s *Selector[T]) GetMulti(ctx *context.Context) ([]*T, error) {
	panic("implement me")
}

// From 指定表名，如果是空字符串，那么将会使用默认表名
func (s *Selector[T]) Build() (*Query, error) {
	var sb strings.Builder

	sb.WriteString("SELECT * FROM ")
	sb.WriteByte('`')

	if s.tbl == "" {
		var t T
		// 提取结构体类型
		typ := reflect.TypeOf(t)

		// 驼峰转下划线
		name := underscoreName(typ.Name())
		sb.WriteString(name)
	} else {
		sb.WriteString(s.tbl)
	}

	sb.WriteByte('`')
	sb.WriteByte(';')

	return &Query{
		SQL: sb.String(),
	}, nil
}


func NewSelector[T any]() *Selector[T] {
	return &Selector[T]{}
}