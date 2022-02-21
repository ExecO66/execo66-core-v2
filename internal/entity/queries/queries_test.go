package queries_test

import (
	"core/internal/config"
	"core/internal/entity"
	"core/internal/entity/enum"
	"core/internal/entity/queries"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	config.Config.Load("../../../config/.env.dev")
	entity.NewDbClient().Connect(config.Config.DbConnString)

	exit := m.Run()

	os.Exit(exit)
}

func TestGetUserByProviderId(t *testing.T) {
	user, err := queries.GetUserByProviderId("198a4d")

	if err != nil {
		t.Fatalf("query error: %v", err)
	}

	assert.Equal(t, "00000000-0000-0000-0000-000000000001", user.Id)
}

func TestInsertUser(t *testing.T) {
	user, err := queries.InsertUser(queries.InsertUserEntity{
		Username:       "test",
		Email:          "test@gmail.com",
		UserStatus:     enum.Student,
		Provider:       enum.Google,
		ProviderId:     "abc123",
		ProfilePicture: "https://picsum.photos/200/200",
	})

	if err != nil {
		t.Fatalf("query error: %v", err)
	}

	assert.Equal(t, "test", user.Username)
	assert.Equal(t, "test@gmail.com", user.Email)
	assert.Equal(t, enum.Student, user.UserStatus)
	assert.Equal(t, enum.Google, user.Provider)
	assert.Equal(t, "abc123", user.ProviderId)
	assert.Equal(t, "https://picsum.photos/200/200", user.ProfilePicture)
}
