package main

import (
	"go-server/app"
	"go-server/config"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Init(config)
	app.Run(":8080")
}
