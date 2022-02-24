package queries

import "core/internal/entity"

func DoesAssignmentExist(assignmentId string) bool {
	sql := `
    SELECT 1 AS exists
    FROM assignment 
    WHERE id=$1;
    `

	var row int
	err := entity.DbClient.Db.QueryRow(sql, assignmentId).Scan(&row)

	if err != nil || row == 0 {
		return false
	} else {
		return true
	}

}
