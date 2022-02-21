package middleware

import (
	"core/internal/session"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		s := session.Default(c)

		user, err := s.GetSessionUser()

		if err != nil {
			c.Status(http.StatusUnauthorized)
			c.Abort()
			return
		}

		c.Set("user", user)
		c.Next()
	}
}
