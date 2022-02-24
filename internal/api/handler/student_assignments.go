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

	assignments, err := queries.GetStudentAssignmentsByUserId(user.Id)

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

var GetStudentAssignmentsById = gin.HandlerFunc(func(c *gin.Context) {
	type Submission struct {
		Id             string `json:"id"`
		SubmitDate     string `json:"submitDate"`
		TestRuns       int    `json:"testRuns"`
		CorrectOutputs int    `json:"correctOuputs"`
	}

	type Assignment struct {
		Id          string       `json:"id"`
		Title       string       `json:"title"`
		Description string       `json:"description"`
		DueDate     string       `json:"dueDate"`
		Submissions []Submission `json:"submissions"`
	}

	id := c.Param("id")

	if id == "" {
		c.Status(http.StatusBadRequest)
		return
	}

	as, err := queries.GetStudentAssignmentsByAssignmentId(id)

	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}

	jsonSubmissions := []Submission{}

AssignmentLoop:
	for _, a := range as {

		if a.SubmissionId == nil {
			continue
		}

		for _, s := range jsonSubmissions {
			if s.Id == *a.SubmissionId { // already added submission, skip rest of loop
				continue AssignmentLoop
			}
		}

		jsonSubmissions = append(
			jsonSubmissions,
			Submission{
				Id:             *a.SubmissionId,
				SubmitDate:     a.SubmitDate,
				TestRuns:       a.SubmissionTestRuns,
				CorrectOutputs: a.SubmissionCorrectOutputs,
			})

	}

	jsonAssignment := Assignment{
		Id:          as[0].AssignmentId,
		Title:       as[0].Title,
		Description: as[0].Description,
		DueDate:     as[0].DueDate,
		Submissions: jsonSubmissions,
	}

	c.JSON(http.StatusOK, jsonAssignment)

})

var PostStudentAssignment = gin.HandlerFunc(func(c *gin.Context) {

	type Body struct {
		AssignmentId string `json:"assignmentId" binding:"required"`
	}

	var body Body
	if err := c.BindJSON(&body); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	user := c.MustGet("user").(*session.SessionUser)

	assignmentExists := queries.DoesAssignmentExist(body.AssignmentId)

	if !assignmentExists {
		c.Status(http.StatusNotFound)
		return
	}

	exists := queries.DoesStudentAssignmentExist(user.Id, body.AssignmentId)

	if exists {
		c.Header("Location", "/student-assignment/"+body.AssignmentId)
		c.Status(http.StatusConflict)
		return
	} else {
		err := queries.InsertStudentAssignment(user.Id, body.AssignmentId)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return
		}
	}

	c.Header("Location", "/student-assignment/"+body.AssignmentId)
	c.Status(http.StatusCreated)
})
