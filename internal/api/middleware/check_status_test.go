package middleware_test

import (
	"core/internal/api/middleware"
	"core/internal/entity/enum"
	"core/internal/session"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestCheckStatus(t *testing.T) {
	pass := false
	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)

	r.Use(func(c *gin.Context) {
		c.Set("user", session.SessionUser{UserStatus: enum.Student})
	})
	r.Use(middleware.CheckStatus(enum.Student))
	r.GET("/ping", func(c *gin.Context) {
		pass = true
	})

	req, _ := http.NewRequest("GET", "/ping", nil)
	r.ServeHTTP(w, req)

	if !pass {
		t.FailNow()
	}
}
