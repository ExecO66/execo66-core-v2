package middleware_test

import (
	"core/internal/api/middleware"
	"core/internal/entity/enum"
	"core/internal/session"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

type mockSession struct {
	hasValidSessionUser bool
}

func (*mockSession) Save()                                {}
func (*mockSession) SetSessionUser(u session.SessionUser) {}
func (m *mockSession) GetSessionUser() (*session.SessionUser, error) {
	if m.hasValidSessionUser {
		return &session.SessionUser{}, nil

	}
	return &session.SessionUser{}, errors.New("failed to get session user")
}
func (*mockSession) SetPendingAuth(s session.PendingAuthState) {}
func (*mockSession) GetPendingAuth() (*session.PendingAuthState, error) {
	return &session.PendingAuthState{}, nil
}
func (*mockSession) RemovePendingAuth() {}

func TestCheckAuthAbort(t *testing.T) {

	session.Default = func(c *gin.Context) session.Session {
		return &mockSession{hasValidSessionUser: false}
	}

	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)

	r.Use(func(c *gin.Context) {
		c.Set("user", session.SessionUser{UserStatus: enum.Student})
	})

	pass := false
	r.Use(middleware.CheckAuth())
	r.GET("/ping", func(c *gin.Context) {
		pass = true
	})

	req, _ := http.NewRequest("GET", "/ping", nil)
	r.ServeHTTP(w, req)

	if pass {
		t.FailNow()
	}
}

func TestCheckAuthPass(t *testing.T) {

	session.Default = func(c *gin.Context) session.Session {
		return &mockSession{hasValidSessionUser: true}
	}

	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)

	r.Use(func(c *gin.Context) {
		c.Set("user", session.SessionUser{UserStatus: enum.Student})
	})

	pass := false
	r.Use(middleware.CheckAuth())
	r.GET("/ping", func(c *gin.Context) {
		pass = true
	})

	req, _ := http.NewRequest("GET", "/ping", nil)
	r.ServeHTTP(w, req)

	if !pass {
		t.FailNow()
	}
}
