package queries

import (
	"core/internal/entity"

	uuid "github.com/satori/go.uuid"
)

type studentAssignment struct {
	AssignmentId             string
	Title                    string
	Description              string
	DueDate                  string
	SubmissionId             *string
	SubmitDate               string
	SubmissionCorrectOutputs int
	SubmissionTestRuns       int
	AvailableUntil           string
}

func GetStudentAssignmentsByAssignmentId(aId string) ([]studentAssignment, error) {
	sql := `
    SELECT 
		a.id AS assignment_id, 
		a.title, 
		a.description, 
		a.due_date, 
		s.id AS submission_id, 
		s.submit_date, 
		s.correct_outputs, 
		s.tests_run,
		a.available_until
	FROM assignment AS a 
	LEFT JOIN submission AS s ON s.assignment_id = a.id 
	WHERE a.id=$1;
    `

	rows, err := entity.DbClient.Db.Query(sql, aId)

	if err != nil {
		return []studentAssignment{}, err
	}

	type dbModel struct {
		AssignmentId uuid.UUID
		SubmissionId uuid.NullUUID
		studentAssignment
	}

	var sa []studentAssignment

	for rows.Next() {
		var s dbModel
		rows.Scan(&s.AssignmentId, &s.Title, &s.Description, &s.DueDate, &s.SubmissionId, &s.SubmitDate, &s.SubmissionCorrectOutputs, &s.SubmissionTestRuns, &s.AvailableUntil)

		var sId *string = nil

		if s.SubmissionId.Valid {
			x := s.SubmissionId.UUID.String()
			sId = &x
		}

		sa = append(sa, studentAssignment{
			AssignmentId:             s.AssignmentId.String(),
			Title:                    s.Title,
			Description:              s.Description,
			DueDate:                  s.DueDate,
			SubmissionId:             sId,
			SubmitDate:               s.SubmitDate,
			SubmissionCorrectOutputs: s.SubmissionCorrectOutputs,
			SubmissionTestRuns:       s.SubmissionTestRuns,
			AvailableUntil:           s.AvailableUntil,
		})
	}

	return sa, err
}
