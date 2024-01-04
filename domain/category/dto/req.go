package dto

type TCreateCategoryRequest struct {
	Name        string `form:"name" validate:"required"`
	Description string `form:"description" validate:"required"`
	Image       string `form:"image"`
}

type TUpdateCategoryRequest struct {
	Name        string `form:"name"`
	Description string `form:"description"`
	Image       string `form:"image"`
}
