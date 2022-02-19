package config

import (
	"os"

	"github.com/joho/godotenv"
)

var IsProd = os.Getenv("APP_ENV") == "production"

var Config = config{}

type config struct {
	DbConnString string
	Port         string
	CookieSecret string
}

func (ev *config) Load() {
	load()
	ev.DbConnString = os.Getenv("PSQL_CONN_STRING")
	ev.Port = os.Getenv("PORT")
	ev.CookieSecret = os.Getenv("COOKIE_SECRET")
}

func load() {
	path := "./config/.env.dev"
	if IsProd {
		path = "./config/.env.prod"
	}
	godotenv.Load(path)
}
