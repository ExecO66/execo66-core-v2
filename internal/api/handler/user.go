package handler

import (
	"core/internal/entity/enum"
	"core/internal/session"
	"net/http"

	"github.com/gin-gonic/gin"
)

var GetUserData = gin.HandlerFunc(func(c *gin.Context) {
	type User struct {
		Id             string          `json:"id"`
		Username       string          `json:"username"`
		Email          string          `json:"email"`
		UserStatus     enum.UserStatus `json:"userStatus"`
		ProfilePicture string          `json:"profilePicture"`
	}

	sessionUser := c.MustGet("user").(*session.SessionUser)

	user := User{
		Id:             sessionUser.Id,
		Username:       sessionUser.Username,
		Email:          sessionUser.Email,
		UserStatus:     sessionUser.UserStatus,
		ProfilePicture: sessionUser.ProfilePicture,
	}

	c.JSON(http.StatusOK, user)
})
