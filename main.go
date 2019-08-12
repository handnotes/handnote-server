package main

import (
	"e.coding.net/handnote/handnote/library"
	"e.coding.net/handnote/handnote/routes"
)

func main() {
	router := routes.SetupRouter()
	router.Run(":" + library.App.HTTPPort)
}
