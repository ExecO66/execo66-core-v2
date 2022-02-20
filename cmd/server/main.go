package main

import (
	"core/internal/api"
	"core/internal/config"
	"core/internal/entity"
)

func main() {
	config.Config.Load(func() string {
		if config.IsProd {
			return "./config/.env.prod"
		}
		return "./config/.env.dev"
	}())

	entity.NewDbClient().Connect(config.Config.DbConnString)

	api.Run(config.Config.Port)
}
