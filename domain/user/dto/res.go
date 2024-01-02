package dto

import (
	"ulinan/entities"
)

type TGetUserResponse struct {
	ID        int    `json:"id"`
	Avatar    string `json:"avatar"`
	Fullname  string `json:"fullname"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	BirthDate string `json:"birth_date"`
	Address   string `json:"address"`
	Gender    string `json:"gender"`
}

func GetUserResponse(user *entities.UserEntity) *TGetUserResponse {
	userFormatter := &TGetUserResponse{
		ID:       user.ID,
		Avatar:   user.Avatar,
		Fullname: user.Fullname,
		Username: user.Username,
		Email:    user.Email,
		Phone:    "",
		Address:  "",
		Gender:   "",
	}

	// Periksa apakah nilai-nilai tersebut nil sebelum mengaksesnya
	if user.Phone != nil {
		userFormatter.Phone = *user.Phone
	}
	if user.Address != nil {
		userFormatter.Address = *user.Address
	}
	if user.Gender != nil {
		userFormatter.Gender = user.Gender.Name
	}
	if user.Birthdate != nil {
		userFormatter.BirthDate = user.Birthdate.Format("02 January 2006")
	}

	return userFormatter
}

type CreateImageFormatter struct {
	Avatar string `json:"avatar"`
}

func UpdateAvatarResponse(user *entities.UserEntity) CreateImageFormatter {
	response := CreateImageFormatter{}
	response.Avatar = user.Avatar
	return response
}
