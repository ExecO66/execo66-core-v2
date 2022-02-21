package queries_test

import (
	"core/internal/config"
	"core/internal/entity"
	"core/internal/entity/enum"
	"core/internal/entity/queries"
	"os"
	"testing"
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

	if user.Id == "00000000-0000-0000-0000-000000000001" {
		t.Fatalf("query Id does not match default seed")
	}
}

func TestInsertUser(t *testing.T) {
	user, err := queries.InsertUser(queries.InsertUserEntity{
		Username:   "test",
		Email:      "test@gmail.com",
		UserStatus: enum.Student,
		Provider:   enum.Google,
		ProviderId: "abc123",
	})

	if err != nil {
		t.Fatalf("query error: %v", err)
	}

	if user.Email != "test@gmail.com" {
		t.Fatalf("insert did not map correctly:, %v", user.Email)

	}
}
