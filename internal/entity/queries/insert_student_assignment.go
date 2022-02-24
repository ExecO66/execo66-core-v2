package queries

import "core/internal/entity"

func InsertStudentAssignment(userId string, assignmentId string) error {
	sql := `
    INSERT INTO student_assignment(student_id, assignment_id) 
    VALUES ($1, $2)
    `

	_, err := entity.DbClient.Db.Exec(sql, userId, assignmentId)

	if err != nil {
		return err
	}

	return nil
}
