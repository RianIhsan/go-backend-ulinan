package dto

type TCreateProductRequest struct {
	Name        string `json:"name" validate:"required"`
	Price       int    `json:"price" validate:"required"`
	Description string `json:"description" validate:"required"`
	Address     string `json:"address"`
	CategoryID  int    `json:"category_id" validate:"required"`
}

type CreateProductImage struct {
	ProductID int    `form:"product_id"`
	Image     string `form:"image" validate:"required"`
}

type TUpdateProductRequest struct {
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Description string `json:"description"`
	Address     string `json:"address"`
	CategoryID  int    `json:"category_id"`
}
