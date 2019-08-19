package v1

import (
	"log"
	"net/http"

	"e.coding.net/handnote/handnote/pkg/redis"
	"github.com/gin-gonic/gin"
)

// RedisCount redis 计数器.
func RedisCount(c *gin.Context) {
	count, err := redis.RedisClient.Incr("count").Result()
	if err != nil {
		log.Fatalln(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": count,
	})
}
