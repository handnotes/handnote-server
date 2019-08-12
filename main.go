package main

import (
	"e.coding.net/handnote/handnote/router"
)

func main() {
	router := router.InitRouter()
	router.Run(":9090")
}
