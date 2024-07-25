package cache

import "time"

func init() {

}

type Cfg struct {
	Driver string
	Redis  struct {
		Addr     string
		Password string
		DB       int
	}
}

type Cache interface {
	Get(key string) interface{}
	Set(key string, value interface{}, duration time.Duration)
	Delete(key string)
	Clear()
}
