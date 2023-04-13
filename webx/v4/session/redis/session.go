package redis

import (
	"context"
)

type Session struct {

}

func (s *Session) Get(ctx context.Context, key string) (any, error) {
	panic("")
}

func (s *Session) Set(ctx context.Context, key string, val string) error {
	panic("")
}

func (s *Session) ID() string {
	panic("")
}