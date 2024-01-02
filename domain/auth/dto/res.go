package dto

import "ulinan/entities"

type TLoginResponse struct {
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	Token    string `json:"access_token"`
}

func LoginResponse(user *entities.UserEntity, token string) *TLoginResponse {
	userFormatter := &TLoginResponse{}
	userFormatter.Username = user.Username
	userFormatter.Avatar = user.Avatar
	userFormatter.Email = user.Email
	userFormatter.Role = user.Role
	userFormatter.Token = token

	return userFormatter
}
