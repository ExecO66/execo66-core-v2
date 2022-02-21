package auth

import (
	"core/internal/config"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func GenGoogleAuthUrl() string {
	url := url.URL{
		Scheme: "https",
		Host:   "accounts.google.com",
		Path:   "/o/oauth2/v2/auth",
	}
	q := url.Query()

	scopes := strings.Join([]string{"https://www.googleapis.com/auth/userinfo.profile",
		"https://www.googleapis.com/auth/userinfo.email"}, " ")

	q.Set("scope", scopes)
	q.Set("access_type", "offline")
	q.Set("include_granted_scopes", "true")
	q.Set("response_type", "code")
	q.Set("redirect_uri", config.Config.GoogleAuthRedirectUri)
	q.Set("client_id", config.Config.GoogleAuthClientId)
	url.RawQuery = q.Encode()

	return url.String()
}

type accessTokenResp struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	TokenType    string `json:"token_type"`
	RefreshToken string `json:"refresh_token"`
}

func GetGoogleAccessTokens(code string) (string, string, error) {
	url := url.URL{
		Scheme: "https",
		Host:   "oauth2.googleapis.com",
		Path:   "/token",
	}

	url.RawQuery = fmt.Sprintf("code=%v&client_id=%v&client_secret=%v&redirect_uri=%v&grant_type=authorization_code", code, config.Config.GoogleAuthClientId, config.Config.GoogleAuthClientSecret, config.Config.GoogleAuthRedirectUri)

	u := url.String()

	resp, err := http.Post(u, "application/x-www-form-urlencoded", nil)

	if err != nil {
		return "", "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", "", errors.New("failed to fetch google access tokens")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", "", err
	}

	var respBody accessTokenResp
	json.Unmarshal(body, &respBody)

	return respBody.AccessToken, respBody.RefreshToken, nil
}

type GoogleUserProfile struct {
	ProviderId string `json:"id"`
	Email      string `json:"email"`
	Name       string `json:"name"`
	ImageUrl   string `json:"picture"`
}

func GetGoogleUserProfile(accessToken string) (GoogleUserProfile, error) {
	url := url.URL{
		Scheme: "https",
		Host:   "www.googleapis.com",
		Path:   "/oauth2/v1/userinfo",
	}

	q := url.Query()
	q.Set("alt", "json")
	q.Set("access_token", accessToken)

	url.RawQuery = q.Encode()

	resp, err := http.Get(url.String())

	if err != nil {
		return GoogleUserProfile{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return GoogleUserProfile{}, errors.New("failed to fetch google user profile")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return GoogleUserProfile{}, err
	}

	var respBody GoogleUserProfile
	json.Unmarshal(body, &respBody)

	return respBody, nil
}
