package entities

import "time"

type ProductEntity struct {
	ID            int                   `gorm:"column:id;primaryKey" json:"id"`
	Name          string                `gorm:"column:name" json:"name"`
	Price         int                   `gorm:"column:price" json:"price"`
	Description   string                `gorm:"column:description" json:"description"`
	Address       string                `gorm:"column:address" json:"address"`
	CategoryID    int                   `gorm:"column:category_id" json:"category_id"`
	Category      CategoryEntity        `gorm:"foreignKey:CategoryID" json:"category"`
	ProductPhotos []ProductPhotosEntity `gorm:"foreignKey:ProductID" json:"product_photos"`
	CreatedAt     time.Time             `gorm:"column:created_at;type:TIMESTAMP" json:"created_at"`
	UpdatedAt     time.Time             `gorm:"column:updated_at;type:TIMESTAMP" json:"updated_at"`
	DeletedAt     *time.Time            `gorm:"column:deleted_at;type:TIMESTAMP NULL;index" json:"deleted_at"`
}

type ProductPhotosEntity struct {
	ID        int        `gorm:"column:id;primaryKey" json:"id"`
	ProductID int        `gorm:"column:product_id" json:"product_id"`
	ImageURL  string     `gorm:"column:image_url" json:"image_url"`
	CreatedAt time.Time  `gorm:"column:created_at;type:TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at;type:TIMESTAMP" json:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at;type:TIMESTAMP NULL;index" json:"deleted_at"`
}

func (ProductEntity) TableName() string {
	return "products"
}

func (ProductPhotosEntity) TableName() string {
	return "product_photos"
}
