package api

import (
	"core/internal/api/handler"
	"core/internal/config"
	"core/internal/entity"
	"core/internal/session"
	"encoding/gob"
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

	gob.Register(&session.PendingAuthState{})
	gob.Register(&session.SessionUser{})
	r.Use(sessions.Sessions("sid", store))

	r.GET("/auth/google", handler.GetGoogleAuth)
	r.GET("/auth/google/callback", handler.GetGoogleAuthCb)
	r.POST("/auth/user-status", handler.PostUserStatus)
	r.Run()
}
