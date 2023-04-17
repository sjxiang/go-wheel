package cache

import (
	"context"
	"sync"
	"time"

	"github.com/sjxiang/go-wheel/cache/internal/errs"
)


type LocalCache struct {
	m sync.Map
}



func (l *LocalCache) Get(ctx context.Context, key string) (any, error) {
	val, ok := l.m.Load(key)
	if !ok {
		return nil, errs.NewKeyNotFound(key)
	}

	return val, nil
}
func (l *LocalCache) Set(ctx context.Context, key string, val string, expiration time.Duration) error {
	l.m.Store(key, val)
	return nil
}

func (l *LocalCache) Delete(ctx context.Context, key string) error {
	l.m.Delete(key)
	return nil
}