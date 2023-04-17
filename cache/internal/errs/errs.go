package errs

import (
	"fmt"
)

func NewKeyNotFound(key string) error {
	return fmt.Errorf("cache: 找不到 key %s", key)
}