package queries

import "core/internal/entity"

func DoesStudentAssignmentExist(userId string, assignmentId string) bool {
	sql := `
    SELECT 1 AS exists
    FROM student_assignment 
    WHERE student_id=$1 AND assignment_id=$2;
    `

	var row int
	err := entity.DbClient.Db.QueryRow(sql, userId, assignmentId).Scan(&row)

	if err != nil || row == 0 {
		return false
	} else {
		return true
	}

}
