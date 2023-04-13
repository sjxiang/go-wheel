package memory

import (
	"context"
	"errors"
	"sync"
	"time"
	"fmt"

	cache "github.com/patrickmn/go-cache"
	"github.com/sjxiang/go-wheel/webx/v4/session"
)

type Session struct {
	id     string
	values sync.Map
}

func (s *Session) Get(ctx context.Context, key string) (string, error) {
	val, ok := s.values.Load(key)
	if !ok {
		return "", errors.New("找不到 key")
	}
	return val.(string), nil
}

func (s *Session) Set(ctx context.Context, key string, val string) error {
	s.values.Store(key, val)
	return nil
}

func (s *Session) ID() string {
	return s.id
}


type Store struct {
	mu         sync.Mutex
	expiration time.Duration
	// 内存缓存
	sessions   *cache.Cache
}


func NewStore(expiration time.Duration) *Store {
	return &Store{
		// 过期时间 + 控制过期检测的间隔
		sessions: cache.New(expiration, time.Second),
	}
}

func (s *Store) Generate(ctx context.Context, id string) (session.Session, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	
	sess :=  &Session{
		id: id,
		values: sync.Map{},
	}
	s.sessions.Set(id, sess, s.expiration)
	return sess, nil
}

func (s *Store) Remove(ctx context.Context, id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	
	s.sessions.Delete(id)
	return nil 
}

func (s *Store) Get(ctx context.Context, id string) (session.Session, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	
	sess, ok := s.sessions.Get(id)
	if !ok {
		return nil, errors.New("找不到 session")
	}
	return sess.(*Session), nil
}


func (s *Store) Refresh(ctx context.Context, id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	
	val, ok := s.sessions.Get(id)
	if !ok {
		return fmt.Errorf("该 id 对应的session 不存在 %s", id)
	}
	s.sessions.Set(id, val, s.expiration)
	return nil
}