package orm

import (
	"context"
	"database/sql"
)

// SELECT 语句
type Querier[T any] interface {
	Get(ctx *context.Context) (*T, error)
	// GetBatch
	GetMulti(ctx *context.Context) ([]*T, error)
}


// UPDATE、DELETE、INSERT 语句
type Executor interface {
	Exec(ctx *context.Context) (sql.Result, error)
}


type QueryBuilder interface {
	Build() (*Query, error)
}


type Query struct {
	SQL  string
	Args []any
}


