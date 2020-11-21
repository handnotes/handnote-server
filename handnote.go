package main

import (
	"fmt"

	"github.com/handnotes/handnote-server/pkg/setting"
	"github.com/handnotes/handnote-server/routes"
)

func main() {
	listenAddr := fmt.Sprintf("0.0.0.0:%d", setting.Server.HTTPPort)
	router := routes.SetupRouter()
	router.Run(listenAddr)
}
