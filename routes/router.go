package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	apiV1 "github.com/handnotes/handnote-server/api/v1"
	"github.com/handnotes/handnote-server/middleware/jwt"
)

func SetupRouter() *gin.Engine {
	router := gin.New()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Pong!"})
	})

	v1 := router.Group("/api/v1")
	{
		v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

		v1.POST("/auth/sendEmail", apiV1.SendEmail)
		v1.POST("/auth/login", apiV1.Login)
		v1.POST("/auth/register", apiV1.Register)

		v1.Use(jwt.JWT())
		{
			v1.PUT("/users/:id", apiV1.UpdateUser)
			v1.GET("/memos", apiV1.ListMemo)
			v1.PUT("/memos/:id", apiV1.UpdateMemo)
		}
	}

	return router
}
