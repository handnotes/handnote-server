package routes

import (
	apiV1 "e.coding.net/handnote/handnote/api/v1"
	"github.com/gin-gonic/gin"
)

// SetupRouter 设置路由.
func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/hello", apiV1.Hello)

	v1 := router.Group("/v1")
	{
		v1.GET("/users", apiV1.GetUsers)
		v1.GET("/count", apiV1.RedisCount)
	}

	return router
}
