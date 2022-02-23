package middleware

import (
	"core/internal/entity/enum"
	"core/internal/session"
	"net/http"

	"github.com/gin-gonic/gin"
)

// assumes user is authorized
func CheckStatus(status enum.UserStatus) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.MustGet("user").(*session.SessionUser)

		if user.UserStatus != status {
			c.Status(http.StatusForbidden)
			c.Abort()
			return
		}

		c.Next()
	}
}
