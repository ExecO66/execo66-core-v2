package main

import (
	"core/internal/api"
	"core/internal/config"
	"core/internal/entity"
)

func main() {
	config.Config.Load()

	entity.NewDbClient().Connect(config.Config.DbConnString)

	api.Run(config.Config.Port)
}
