package utils

import (
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
)

func RedisScan(cacheKey string) (keys []string, err error) {
	var cursor uint64
	for {
		var newKeys []string
		cursor, newKeys, err = g.Redis().Scan(ctx, cursor, gredis.ScanOption{
			Match: cacheKey,
			Count: 1000,
		})
		if err != nil {
			g.Log().Errorf(ctx, "Scan failed: %v", err)
			break
		}
		keys = append(keys, newKeys...)
		if cursor == 0 {
			break
		}
	}
	return
}
