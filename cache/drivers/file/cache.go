package file

import (
	"github.com/ayflying/yunloli-go/cache"
	"time"
)

type cacheMod struct {
}

func (c *cacheMod) Get(key string) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c *cacheMod) Set(key string, value interface{}, duration time.Duration) {
	//TODO implement me
	panic("implement me")
}

func (c *cacheMod) Delete(key string) {
	//TODO implement me
	panic("implement me")
}

func (c *cacheMod) Clear() {
	//TODO implement me
	panic("implement me")
}

func New() cache.Cache {
	var c = &cacheMod{}
	return c
}
