package api

import (
	"core/internal/config"
	"core/internal/entity"
	"log"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/postgres"
	"github.com/gin-gonic/gin"
)

var r = gin.New()

func Run(port string) {
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	store, err := postgres.NewStore(entity.DbClient.Db, []byte(config.Config.CookieSecret))

	if err != nil {
		log.Fatal("error creating postgres store")
	}

	r.Use(sessions.Sessions("sid", store))

	r.Run()
}
