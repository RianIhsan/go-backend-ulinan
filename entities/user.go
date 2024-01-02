package entities

import "time"

type UserEntity struct {
	ID        int           `gorm:"column:id;primaryKey" json:"id"`
	Fullname  string        `gorm:"column:fullname" json:"fullname"`
	Username  string        `gorm:"column:username" json:"username"`
	Avatar    string        `gorm:"column:avatar" json:"avatar"`
	Email     string        `gorm:"column:email" json:"email"`
	Password  string        `gorm:"column:password" json:"password"`
	Role      string        `gorm:"column:role" json:"role"`
	Phone     *string       `gorm:"column:phone;default:null" json:"phone,omitempty"`
	Address   *string       `gorm:"column:address;default:null" json:"address,omitempty"`
	GenderID  *int          `gorm:"column:gender_id;default:null" json:"gender_id,omitempty"`
	Gender    *GenderEntity `gorm:"foreignKey:GenderID" json:"gender,omitempty"`
	Birthdate *time.Time    `gorm:"column:birthdate;default:null" json:"birthdate,omitempty"`
	CreatedAt time.Time     `gorm:"column:created_at;type:TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time     `gorm:"column:updated_at;type:TIMESTAMP" json:"updated_at"`
	DeletedAt *time.Time    `gorm:"column:deleted_at;type:TIMESTAMP NULL;index" json:"deleted_at"`
}

type GenderEntity struct {
	ID   int    `gorm:"column:id;primaryKey" json:"id"`
	Name string `gorm:"column:name" json:"name"`
}

func (UserEntity) TableName() string {
	return "users"
}
func (GenderEntity) TableName() string {
	return "genders"
}
