package router

import (
	apiV1 "e.coding.net/handnote/handnote/api/v1"
	"github.com/gin-gonic/gin"
)

// InitRouter 初始化路由.
func InitRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/hello", apiV1.Hello)

	return router
}
