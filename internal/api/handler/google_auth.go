package handler

import (
	"core/internal/auth"
	"core/internal/config"
	"core/internal/entity/enum"
	"core/internal/entity/queries"
	"core/internal/session"
	"net/http"

	"github.com/gin-gonic/gin"
)

var GetGoogleAuth = gin.HandlerFunc(func(c *gin.Context) {
	c.Redirect(http.StatusPermanentRedirect, auth.GenGoogleAuthUrl())
})

var GetGoogleAuthCb = gin.HandlerFunc(func(c *gin.Context) {
	q := c.Request.URL.Query()
	code := q.Get("code")
	error := q.Get("error")

	if code == "" || error != "" {
		c.Status(http.StatusBadRequest)
		return
	}

	accessToken, _, err := auth.GetGoogleAccessTokens(code)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to fetch access tokens"})
		return
	}

	googleUser, err := auth.GetGoogleUserProfile(accessToken)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to fetch google data"})
		return
	}

	user, queryErr := queries.GetUserByProviderId(googleUser.ProviderId)

	s := session.Default(c)

	// user does not exist yets
	if queryErr != nil {
		s.SetPendingAuth(session.PendingAuthState{
			Username:       googleUser.Name,
			Email:          googleUser.Email,
			Provider:       enum.Google,
			ProviderId:     googleUser.ProviderId,
			ProfilePicture: googleUser.ImageUrl,
		})
		s.Save()
		c.Redirect(http.StatusPermanentRedirect, config.Config.ClientBaseUrl+"/auth/user-status")
		return
	}

	s.SetSessionUser(session.SessionUser(user))
	s.Save()

	c.Redirect(http.StatusPermanentRedirect, config.Config.ClientBaseUrl+"/dash")
})
