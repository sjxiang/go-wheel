package orm

import "strings"

type SelectBuilder struct {
	builder   *strings.Builder
	column    []string
	tableName string
	where     []func(s *SelectBuilder)
	args      []interface{}
	orderby   string
	offset    *int64
	limit     *int64
}

func (s *SelectBuilder) Select(field ...string) *SelectBuilder {
	s.column = append(s.column, field...)
	return s
}