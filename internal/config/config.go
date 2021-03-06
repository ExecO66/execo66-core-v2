package config

import (
	"os"

	"github.com/joho/godotenv"
)

var IsProd = os.Getenv("APP_ENV") == "production"

var Config = config{}

type config struct {
	Port          string
	CookieSecret  string
	ClientBaseUrl string
	GoogleAuth
	DbConfig
}

func (ev *config) Load(path string) {
	godotenv.Load(path)

	ev.DbConnString = os.Getenv("PSQL_CONN_STRING")
	ev.TestDbConnString = os.Getenv("TEST_PSQL_CONN_STRING")
	ev.Port = os.Getenv("PORT")
	ev.CookieSecret = os.Getenv("COOKIE_SECRET")
	ev.GoogleAuthClientId = os.Getenv("GOOGLE_AUTH_CLIENT_ID")
	ev.GoogleAuthClientSecret = os.Getenv("GOOGLE_AUTH_CLIENT_SECRET")
	ev.GoogleAuthRedirectUri = os.Getenv("GOOGLE_AUTH_REDIRECT_URI")
	ev.ClientBaseUrl = os.Getenv("CLIENT_BASE_URL")
}
