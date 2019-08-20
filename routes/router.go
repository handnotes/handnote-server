package routes

import (
	apiV1 "e.coding.net/handnote/handnote/api/v1"
	"github.com/gin-gonic/gin"
)

// SetupRouter 设置路由.
func SetupRouter() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		v1.GET("/users", apiV1.GetUsers)
		v1.POST("/users", apiV1.CreateUser)
		v1.POST("/sendEmail", apiV1.SendEmail)
	}

	return router
}
