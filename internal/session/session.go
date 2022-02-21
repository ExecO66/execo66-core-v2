package session

import (
	"core/internal/entity/enum"
	"errors"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

var (
	sessionUserKey = "user"
	pendingAuthKey = "pending-auth"
)

type Session struct {
	session sessions.Session
}

func Default(c *gin.Context) *Session {
	session := sessions.Default(c)
	return &Session{session}
}

type SessionUser struct {
	Id         string
	Username   string
	Email      string
	UserStatus enum.UserStatus
	Provider   enum.LoginProvider
	ProviderId string
}

func (s *Session) Save() {
	s.session.Save()
}

func (s *Session) SetSessionUser(sessionUser SessionUser) {
	s.session.Set(sessionUserKey, sessionUser)
}

func (s *Session) GetSessionUser() (*SessionUser, error) {
	state := s.session.Get(sessionUserKey)
	if v, ok := state.(*SessionUser); ok {
		return v, nil
	}
	return &SessionUser{}, errors.New("unable to get session user")
}

type PendingAuthState struct {
	Username   string
	Email      string
	Provider   enum.LoginProvider
	ProviderId string
}

func (s *Session) SetPendingAuth(state PendingAuthState) {
	s.session.Set(pendingAuthKey, state)
}

func (s *Session) GetPendingAuth() (*PendingAuthState, error) {
	state := s.session.Get(pendingAuthKey)
	if v, ok := state.(*PendingAuthState); ok {
		return v, nil
	}

	return &PendingAuthState{}, errors.New("unable to get pending auth state")
}

func (s *Session) RemovePendingAuth() {
	s.session.Delete(pendingAuthKey)
}
