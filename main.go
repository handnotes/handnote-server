package main

import (
	"e.coding.net/handnote/handnote/routes"
)

func main() {
	router := routes.SetupRouter()
	router.Run(":9090")
}
