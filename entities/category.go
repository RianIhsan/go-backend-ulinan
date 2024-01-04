package entities

import "time"

type CategoryEntity struct {
	ID          int        `gorm:"column:id;primaryKey" json:"id"`
	Name        string     `gorm:"column:name" json:"name"`
	Description string     `gorm:"column:description" json:"description"`
	Image       string     `gorm:"image" json:"image"`
	CreatedAt   time.Time  `gorm:"column:created_at;type:TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"column:updated_at;type:TIMESTAMP" json:"updated_at"`
	DeletedAt   *time.Time `gorm:"column:deleted_at;type:TIMESTAMP NULL;index" json:"deleted_at"`
}

func (CategoryEntity) TableName() string {
	return "categories"
}
