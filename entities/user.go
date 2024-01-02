package entities

import "time"

type UserEntity struct {
	ID        int        `gorm:"column:id;primaryKey" json:"id"`
	Fullname  string     `gorm:"column:fullname" json:"fullname"`
	Username  string     `gorm:"column:username" json:"username"`
	Avatar    string     `gorm:"column:avatar" json:"avatar"`
	Email     string     `gorm:"column:email" json:"email"`
	Password  string     `gorm:"column:password" json:"password"`
	Role      string     `gorm:"column:role" json:"role"`
	CreatedAt time.Time  `gorm:"column:created_at;type:TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at;type:TIMESTAMP" json:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at;type:TIMESTAMP NULL;index" json:"deleted_at"`
}

func (UserEntity) TableName() string {
	return "users"
}
