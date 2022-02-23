package handler

import (
	"core/internal/entity/queries"
	"core/internal/session"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
   - auth
   - student
*/

var GetAllStudentAssignment = gin.HandlerFunc(func(c *gin.Context) {
	type TeacherInfo struct {
		Username       string `json:"username"`
		ProfilePicture string `json:"profilePicture"`
	}

	type StudentAssignment struct {
		AssignmentId       string      `json:"assignmentId"`
		Title              string      `json:"title"`
		Description        string      `json:"description"`
		DueDate            string      `json:"dueDate"`
		RecentSubmissionId *string     `json:"recentSubmissionId,omitempty"`
		TeacherInfo        TeacherInfo `json:"teacherInfo"`
	}

	user := c.MustGet("user").(*session.SessionUser)

	assignments, err := queries.GetStudentAssignmentsById(user.Id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to retrieve user assignments"})
		return
	}

	var jsonAssignments []StudentAssignment
	for _, a := range assignments {
		jsonAssignments = append(jsonAssignments,
			StudentAssignment{
				AssignmentId:       a.AssignmentId,
				Title:              a.Title,
				Description:        a.Description,
				DueDate:            a.DueDate,
				RecentSubmissionId: a.RecentSubmissionId,
				TeacherInfo: TeacherInfo{
					Username:       a.TeacherUsername,
					ProfilePicture: a.TeacherProfilePicture,
				},
			})
	}

	c.JSON(http.StatusOK, map[string][]StudentAssignment{"assignments": jsonAssignments})
})
