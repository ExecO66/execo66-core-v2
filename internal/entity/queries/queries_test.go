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
	entity.NewDbClient().Connect(config.Config.TestDbConnString)

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

func TestGetUserById(t *testing.T) {
	user, err := queries.GetUserById("00000000-0000-0000-0000-000000000001")

	assert.Nil(t, err)

	assert.Equal(t, "00000000-0000-0000-0000-000000000001", user.Id)
	assert.Equal(t, "Bob Schoolers", user.Username)
	assert.Equal(t, "bob@gmail.com", user.Email)
	assert.Equal(t, "https://picsum.photos/200/200", user.ProfilePicture)
	assert.Equal(t, enum.Google, user.Provider)
	assert.Equal(t, enum.Student, user.UserStatus)
	assert.Equal(t, "198a4d", user.ProviderId)
}

func TestGetStudentAssignmentsByUserId(t *testing.T) {
	assignments, err := queries.GetStudentAssignmentsByUserId("00000000-0000-0000-0000-000000000001")

	assert.Nil(t, err)
	assert.Len(t, assignments, 3)
	assert.Equal(t, "CS A Lab 2", assignments[0].Title)
	assert.Equal(t, "CS A Lab 3", assignments[1].Title)
	assert.Equal(t, "CS A Lab 4", assignments[2].Title)
	assert.Nil(t, assignments[2].RecentSubmissionId)
}

func TestGetStudentAssignmentsByAssignmentId(t *testing.T) {
	as, err := queries.GetStudentAssignmentsByAssignmentId("00000000-0000-0000-0000-000000000001")

	assert.Nil(t, err)
	assert.Len(t, as, 1)
	assert.Nil(t, as[0].SubmissionId)

	as1, err := queries.GetStudentAssignmentsByAssignmentId("00000000-0000-0000-0000-000000000002")

	assert.Nil(t, err)
	assert.Len(t, as1, 1)
	assert.Equal(t, 3, as1[0].SubmissionTestRuns)
	assert.Equal(t, 2, as1[0].SubmissionCorrectOutputs)
}

func TestDoesStudentAssignmentExist(t *testing.T) {
	exists1 := queries.DoesStudentAssignmentExist("00000000-0000-0000-0000-000000000001", "00000000-0000-0000-0000-000000000002")
	exists2 := queries.DoesStudentAssignmentExist("00000000-0000-0000-0000-000000000001", "00000000-0000-0000-0000-000000000001")

	assert.True(t, exists1)
	assert.False(t, exists2)
}

func TestInsertStudentAssignment(t *testing.T) {
	err := queries.InsertStudentAssignment("00000000-0000-0000-0000-000000000001", "00000000-0000-0000-0000-000000000001")

	assert.Nil(t, err)
}

func TestDoesAssignmentExist(t *testing.T) {
	exists := queries.DoesAssignmentExist("00000000-0000-0000-0000-000000000002")
	assert.True(t, exists)
}
