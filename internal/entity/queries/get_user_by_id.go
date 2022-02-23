package queries

import (
	"core/internal/entity"
	"core/internal/entity/enum"

	uuid "github.com/jackc/pgx/pgtype/ext/satori-uuid"
)

type getUserByIdModel struct {
	Id             string
	Username       string
	Email          string
	UserStatus     enum.UserStatus
	Provider       enum.LoginProvider
	ProviderId     string
	ProfilePicture string
}

func GetUserById(userId string) (getUserByIdModel, error) {
	sql := `
	SELECT 
		id, 
		username, 
		email, 
		user_status, 
		provider, 
		provider_id, 
		profile_picture 
	FROM "user" 
	WHERE id=$1;`

	var dbModel struct {
		Id uuid.UUID
		getUserByIdModel
	}

	err := entity.DbClient.Db.QueryRow(sql, userId).Scan(&dbModel.Id, &dbModel.Username, &dbModel.Email, &dbModel.UserStatus, &dbModel.Provider, &dbModel.ProviderId, &dbModel.ProfilePicture)

	if err != nil {
		return getUserByIdModel{}, err
	}

	return getUserByIdModel{
		Id:             dbModel.Id.UUID.String(),
		Username:       dbModel.Username,
		Email:          dbModel.Email,
		UserStatus:     dbModel.UserStatus,
		Provider:       dbModel.Provider,
		ProviderId:     dbModel.ProviderId,
		ProfilePicture: dbModel.ProfilePicture,
	}, err

}
