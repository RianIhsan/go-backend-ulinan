package dto

type TUpdateAvatarRequest struct {
	Avatar string `form:"avatar"`
}

type TUpdateUserRequest struct {
	Fullname  *string                   `json:"fullname"`
	Username  *string                   `json:"username"`
	Email     *string                   `json:"email"`
	Phone     *string                   `json:"phone,omitempty"`
	Birthdate *string                   `json:"birthdate,omitempty"`
	Address   *string                   `json:"address,omitempty"`
	GenderID  *TUpdateUserGenderRequest `json:"gender_id,omitempty"`
}

type TUpdateUserGenderRequest struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type TUpdatePasswordRequest struct {
	OldPassword     string `json:"old_password"`
	NewPassword     string `json:"new_password"`
	ConfirmPassword string `json:"confirm_password"`
}
