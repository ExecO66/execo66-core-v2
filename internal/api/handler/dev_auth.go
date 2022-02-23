package handler

import (
	"core/internal/entity/queries"
	"core/internal/session"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
	- dev only
*/

var GetDevAuth = gin.HandlerFunc(func(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "please request a valid user statu"})
		return
	}

	user, err := queries.GetUserById(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "invalid id"})
		return
	}

	s := session.Default(c)
	s.SetSessionUser(session.SessionUser(user))
	s.Save()
	c.Status(http.StatusOK)
})
