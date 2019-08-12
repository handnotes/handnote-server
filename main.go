package main

import (
	"github.com/gin-gonic/gin"

	api "e.coding.net/handnote/handnote/api/v1"
)

func main() {
	router := gin.Default()
	router.GET("/hello", api.Hello)
	router.Run(":9090")
}
