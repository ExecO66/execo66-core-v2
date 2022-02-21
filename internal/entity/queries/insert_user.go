package queries

import (
	"core/internal/entity"
	"core/internal/entity/enum"
)

type InsertUserEntity struct {
	Username       string
	Email          string
	UserStatus     enum.UserStatus
	Provider       enum.LoginProvider
	ProviderId     string
	ProfilePicture string
}

type InsertUserModel struct {
	Id         string
	Username   string
	Email      string
	UserStatus enum.UserStatus
	Provider   enum.LoginProvider
	ProviderId string
}

func InsertUser(user InsertUserEntity) (InsertUserModel, error) {
	sql := `
    INSERT INTO public.user (
        username, 
        email, 
        user_status, 
        provider, 
        provider_id,
		profile_picture
    ) VALUES ($1, $2, $3, $4, $5, $6) 
    RETURNING id;`

	id := ""
	err := entity.DbClient.Db.QueryRow(sql, user.Username, user.Email, user.UserStatus, user.Provider, user.ProviderId, user.ProfilePicture).Scan(&id)

	return InsertUserModel{Id: id, Username: user.Username, Email: user.Email, UserStatus: user.UserStatus, Provider: user.Provider, ProviderId: user.ProviderId}, err

}
