package session

import (
	"core/internal/entity/enum"
	"errors"

	gsession "github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

var (
	sessionUserKey = "user"
	pendingAuthKey = "pending-auth"
)

type Session interface {
	Save()
	SetSessionUser(user SessionUser)
	GetSessionUser() (*SessionUser, error)
	SetPendingAuth(state PendingAuthState)
	GetPendingAuth() (*PendingAuthState, error)
	RemovePendingAuth()
}

type session struct {
	session gsession.Session
}

var Default = func(c *gin.Context) Session {
	s := gsession.Default(c)
	return &session{s}
}

type SessionUser struct {
	Id             string
	Username       string
	Email          string
	UserStatus     enum.UserStatus
	Provider       enum.LoginProvider
	ProviderId     string
	ProfilePicture string
}

func (s *session) Save() {
	s.session.Save()
}

func (s *session) SetSessionUser(sessionUser SessionUser) {
	s.session.Set(sessionUserKey, sessionUser)
}

func (s *session) GetSessionUser() (*SessionUser, error) {
	state := s.session.Get(sessionUserKey)
	if v, ok := state.(*SessionUser); ok {
		return v, nil
	}
	return &SessionUser{}, errors.New("unable to get session user")
}

type PendingAuthState struct {
	Username       string
	Email          string
	Provider       enum.LoginProvider
	ProviderId     string
	ProfilePicture string
}

func (s *session) SetPendingAuth(state PendingAuthState) {
	s.session.Set(pendingAuthKey, state)
}

func (s *session) GetPendingAuth() (*PendingAuthState, error) {
	state := s.session.Get(pendingAuthKey)
	if v, ok := state.(*PendingAuthState); ok {
		return v, nil
	}

	return &PendingAuthState{}, errors.New("unable to get pending auth state")
}

func (s *session) RemovePendingAuth() {
	s.session.Delete(pendingAuthKey)
}
