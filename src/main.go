package main

import (
	"ctrl/api/router"
	"ctrl/config"
)

func main() {
	config.InitConfig()
	router.SetupRouter()
}
