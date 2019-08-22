package main

import (
	"fmt"

	"e.coding.net/handnote/handnote/pkg/setting"
	"e.coding.net/handnote/handnote/routes"
)

func main() {
	listenAddr := fmt.Sprintf("0.0.0.0:%d", setting.Server.HTTPPort)
	router := routes.SetupRouter()
	router.Run(listenAddr)
}
