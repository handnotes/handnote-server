package main

import (
	"fmt"

	"github.com/handnotes/handnote-server/docs"
	"github.com/handnotes/handnote-server/pkg/setting"
	"github.com/handnotes/handnote-server/routes"
)

func main() {
	docs.SwaggerInfo.Title = "Handnote API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/api/v1"

	listenAddr := fmt.Sprintf("127.0.0.1:%d", setting.Server.HTTPPort)
	router := routes.SetupRouter()

	_ = router.Run(listenAddr)
}
