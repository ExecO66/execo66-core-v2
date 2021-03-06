package queries

import (
	"core/internal/entity"
	"core/internal/entity/enum"

	uuid "github.com/jackc/pgx/pgtype/ext/satori-uuid"
)

type GetUserByProviderIdModel struct {
	Id             string
	Username       string
	Email          string
	UserStatus     enum.UserStatus
	Provider       enum.LoginProvider
	ProviderId     string
	ProfilePicture string
}

func GetUserByProviderId(providerId string) (GetUserByProviderIdModel, error) {
	sql := `
    SELECT 
        id, 
        username, 
        email, 
        user_status, 
        provider, 
        provider_id,
		profile_picture
    FROM public.user 
    WHERE provider_id=$1;`

	var dbModel struct {
		Id uuid.UUID
		GetUserByProviderIdModel
	}

	err := entity.DbClient.Db.QueryRow(sql, providerId).Scan(&dbModel.Id, &dbModel.Username, &dbModel.Email, &dbModel.UserStatus, &dbModel.Provider, &dbModel.ProviderId, &dbModel.ProfilePicture)

	model := GetUserByProviderIdModel{
		Id:             dbModel.Id.UUID.String(),
		Username:       dbModel.Username,
		Email:          dbModel.Email,
		UserStatus:     dbModel.UserStatus,
		Provider:       dbModel.Provider,
		ProviderId:     dbModel.ProviderId,
		ProfilePicture: dbModel.ProfilePicture,
	}

	return model, err
}
