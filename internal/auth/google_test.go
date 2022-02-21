package auth_test

import (
	"core/internal/auth"
	"core/internal/config"
	"testing"
)

func TestGenGoogleAuthUrl(t *testing.T) {
	o1 := config.Config.GoogleAuthRedirectUri
	o2 := config.Config.GoogleAuthClientId

	config.Config.GoogleAuthRedirectUri = "abc"
	config.Config.GoogleAuthClientId = "def"

	u := auth.GenGoogleAuthUrl()

	expected := `https://accounts.google.com/o/oauth2/v2/auth?access_type=offline&client_id=def&include_granted_scopes=true&redirect_uri=abc&response_type=code&scope=https%3A%2F%2Fwww.googleapis.com%2Fauth%2Fuserinfo.profile+https%3A%2F%2Fwww.googleapis.com%2Fauth%2Fuserinfo.email`

	if u != expected {
		t.Fail()
	}

	config.Config.GoogleAuthRedirectUri = o1
	config.Config.GoogleAuthClientId = o2
}
