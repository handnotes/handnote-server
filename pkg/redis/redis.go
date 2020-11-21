package redis

import (
	"github.com/go-redis/redis"
	"github.com/handnotes/handnote-server/pkg/setting"
)

// RedisClient redis 连接对象
var RedisClient *redis.Client

func init() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     setting.Redis.Addr,
		Password: setting.Redis.Password,
		DB:       setting.Redis.DB,
	})
}
