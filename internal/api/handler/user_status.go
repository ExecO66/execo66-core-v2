package handler

import (
	"core/internal/config"
	"core/internal/entity/enum"
	"core/internal/entity/queries"
	"core/internal/session"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type postUserBody struct {
	UserStatus enum.UserStatus
}

func (b *postUserBody) UnmarshalJSON(data []byte) error {
	// avoid recursion
	type Aux postUserBody
	var a *Aux = (*Aux)(b)
	err := json.Unmarshal(data, &a)
	if err != nil {
		return err
	}

	switch b.UserStatus {
	case enum.Student, enum.Teacher:
		return nil
	default:
		b.UserStatus = ""
		return errors.New("invalid enum value")
	}

}

var PostUserStatus = gin.HandlerFunc(func(c *gin.Context) {

	s := session.Default(c)

	state, err := s.GetPendingAuth()

	if err != nil {
		c.Status(http.StatusUnauthorized)
		return
	}

	jsonData, err := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "request must have body"})
		c.Abort()
		return
	}

	var body postUserBody
	json.Unmarshal(jsonData, &body)

	if body.UserStatus == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user status"})
		return
	}

	existingUser, queryErr := queries.GetUserByProviderId(state.ProviderId)

	sessionUser := session.SessionUser(existingUser)

	if queryErr != nil {

		newUser, err := queries.InsertUser(queries.InsertUserEntity{
			Username:       state.Username,
			Email:          state.Email,
			UserStatus:     body.UserStatus,
			Provider:       state.Provider,
			ProviderId:     state.ProviderId,
			ProfilePicture: state.ProfilePicture,
		})

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to add user"})
			return
		}

		sessionUser = session.SessionUser(newUser)

	}

	s.RemovePendingAuth()
	s.SetSessionUser(sessionUser)

	s.Save()

	c.Redirect(http.StatusPermanentRedirect, config.Config.ClientBaseUrl+"/dash")
})
