package queries

import (
	"core/internal/entity"

	uuid "github.com/jackc/pgx/pgtype/ext/satori-uuid"
)

type getUserAssignmentsByIdModel struct {
	AssignmentId          string
	Title                 string
	Description           string
	DueDate               string
	TeacherUsername       string
	TeacherProfilePicture string
	RecentSubmissionId    *string
}

func GetStudentAssignmentsById(userId string) ([]getUserAssignmentsByIdModel, error) {
	sql := `
    SELECT 
        a.id AS assignment_id,
        a.title, 
        a.description, 
        a.due_date, 
        owner.username AS teacher_username, 
        owner.profile_picture as teacher_profile_picture, 
        s.id as recent_submission_id 
    FROM "user" 
    INNER JOIN student_assignment ON "user".id = student_assignment.student_id 
    INNER JOIN assignment AS a ON student_assignment.assignment_id = a.id 
    INNER JOIN "user" AS owner ON a.owner_id = "owner".id 
    LEFT JOIN submission AS s 
    ON 
        s.id = (
            SELECT s1.id
            FROM submission AS s1 
            WHERE a.id = s1.assignment_id 
            ORDER BY s1.submit_date DESC
            LIMIT 1
        )  
    WHERE "user".id=$1;
    `

	rows, err := entity.DbClient.Db.Query(sql, userId)
	if err != nil {
		return make([]getUserAssignmentsByIdModel, 0), err
	}
	defer rows.Close()

	var assignments []getUserAssignmentsByIdModel

	for rows.Next() {
		type dbModel struct {
			AssignmentId uuid.UUID
			getUserAssignmentsByIdModel
		}

		var a dbModel

		if err := rows.Scan(&a.AssignmentId, &a.Title, &a.Description, &a.DueDate, &a.TeacherUsername, &a.TeacherProfilePicture, &a.RecentSubmissionId); err != nil {
			return make([]getUserAssignmentsByIdModel, 0), err
		}

		assignments = append(assignments, getUserAssignmentsByIdModel{
			AssignmentId:          a.AssignmentId.UUID.String(),
			Title:                 a.Title,
			Description:           a.Description,
			DueDate:               a.DueDate,
			TeacherUsername:       a.TeacherUsername,
			TeacherProfilePicture: a.TeacherProfilePicture,
			RecentSubmissionId:    a.RecentSubmissionId,
		})
	}

	return assignments, nil
}
