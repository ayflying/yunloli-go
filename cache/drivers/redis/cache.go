package redis

import (
	"github.com/ayflying/yunloli-go/cache"
	"time"
)

type cacheMod struct{}

func (s *cacheMod) Set(key string, value interface{}, duration time.Duration) {
	//TODO implement me
	panic("implement me")
}

func (s *cacheMod) Delete(key string) {
	//TODO implement me
	panic("implement me")
}

func (s *cacheMod) Clear() {
	//TODO implement me
	panic("implement me")
}

func New() cache.Cache {
	var c = &cacheMod{}
	return c
}

func (s *cacheMod) Get(key string) (res interface{}) {

	return
}
