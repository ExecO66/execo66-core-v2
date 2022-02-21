package auth_test

import (
	"bytes"
	"core/internal/auth"
	"core/internal/config"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenGoogleAuthUrl(t *testing.T) {
	o1 := config.Config.GoogleAuthRedirectUri
	o2 := config.Config.GoogleAuthClientId

	config.Config.GoogleAuthRedirectUri = "abc"
	config.Config.GoogleAuthClientId = "def"

	u := auth.GenGoogleAuthUrl()

	expected := `https://accounts.google.com/o/oauth2/v2/auth?access_type=offline&client_id=def&include_granted_scopes=true&redirect_uri=abc&response_type=code&scope=https%3A%2F%2Fwww.googleapis.com%2Fauth%2Fuserinfo.profile+https%3A%2F%2Fwww.googleapis.com%2Fauth%2Fuserinfo.email`

	assert.Equal(t, expected, u)

	config.Config.GoogleAuthRedirectUri = o1
	config.Config.GoogleAuthClientId = o2
}

func TestGetGoogleAccessTokens(t *testing.T) {
	oReq := auth.PostGoogleTokenRequest
	defer func() { auth.PostGoogleTokenRequest = oReq }()

	json := `{"access_token":"test_ac_token", "expires_in": 123, "token_type":"Bearer", "refresh_token":"test_ref_token"}`

	auth.PostGoogleTokenRequest = func(url string) (*http.Response, error) {
		return &http.Response{Body: ioutil.NopCloser(bytes.NewBufferString(json)), StatusCode: http.StatusOK}, nil
	}

	expectedAc, expectedRef := "test_ac_token", "test_ref_token"

	aToken, rToken, err := auth.GetGoogleAccessTokens("code")

	assert.Nil(t, err)
	assert.Equal(t, expectedAc, aToken)
	assert.Equal(t, expectedRef, rToken)
}

func TestGetGoogleUserProfile(t *testing.T) {
	oReq := auth.GetGoogleUserRequest
	defer func() { auth.GetGoogleUserRequest = oReq }()

	json := `{"id":"test_id", "email":"test_email", "name":"test_name", "picture":"image_url"}`

	auth.GetGoogleUserRequest = func(url string) (*http.Response, error) {
		return &http.Response{Body: ioutil.NopCloser(bytes.NewBufferString(json)), StatusCode: http.StatusOK}, nil
	}

	resp, err := auth.GetGoogleUserProfile("access_token")

	assert.Nil(t, err)
	assert.Equal(t, "test_email", resp.Email)
	assert.Equal(t, "test_id", resp.ProviderId)
	assert.Equal(t, "test_name", resp.Name)
	assert.Equal(t, "image_url", resp.ImageUrl)

}
