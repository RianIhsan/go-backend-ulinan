package dto

import "ulinan/entities"

type TCreateCategoryResponse struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

func CreateCategoryResponse(category *entities.CategoryEntity) TCreateCategoryResponse {
	response := TCreateCategoryResponse{}
	response.Id = category.ID
	response.Name = category.Name
	response.Description = category.Description
	response.Image = category.Image
	return response
}

func GetCategoryResponse(category *entities.CategoryEntity) TCreateCategoryResponse {
	response := TCreateCategoryResponse{}
	response.Id = category.ID
	response.Name = category.Name
	response.Description = category.Description
	response.Image = category.Image
	return response
}

func GetPaginationCategories(categories []*entities.CategoryEntity) []*TCreateCategoryResponse {
	var categoryFormatters []*TCreateCategoryResponse

	for _, category := range categories {
		formattedCategory := GetCategoryResponse(category)
		categoryFormatters = append(categoryFormatters, &formattedCategory)
	}

	return categoryFormatters
}
