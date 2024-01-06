package dto

import "ulinan/entities"

type TCreateProductResponse struct {
	ID          int                     `json:"id"`
	Name        string                  `json:"name"`
	Price       int                     `json:"price"`
	Description string                  `json:"description"`
	Address     string                  `json:"address"`
	Category    string                  `json:"category"`
	Image       []ProductImageFormatter `json:"image"`
}

type ProductImageFormatter struct {
	ID  int    `json:"id"`
	URL string `json:"image_url"`
}

func GetProductById(product *entities.ProductEntity) TCreateProductResponse {
	response := TCreateProductResponse{}
	response.ID = product.ID
	response.Name = product.Name
	response.Price = product.Price
	response.Description = product.Description
	response.Address = product.Address
	response.Category = product.Category.Name

	var productImages []ProductImageFormatter
	for _, productImage := range product.ProductPhotos {
		if productImage.DeletedAt == nil {
			image := ProductImageFormatter{
				ID:  productImage.ID,
				URL: productImage.ImageURL,
			}
			productImages = append(productImages, image)
		}
	}

	response.Image = productImages
	return response
}

type TGetAllProductsResponse struct {
	ID          int                     `json:"id"`
	Name        string                  `json:"name"`
	Price       int                     `json:"price"`
	Category    string                  `json:"category"`
	Address     string                  `json:"address"`
	Description string                  `json:"description"`
	Image       []ProductImageFormatter `json:"image"`
}

func GetAllProductsResponse(product *entities.ProductEntity) TGetAllProductsResponse {
	response := TGetAllProductsResponse{}
	response.ID = product.ID
	response.Name = product.Name
	response.Price = product.Price
	response.Description = product.Description
	response.Address = product.Address
	response.Category = product.Category.Name

	var productImages []ProductImageFormatter
	for _, productImage := range product.ProductPhotos {
		if productImage.DeletedAt == nil {
			image := ProductImageFormatter{
				ID:  productImage.ID,
				URL: productImage.ImageURL,
			}
			productImages = append(productImages, image)
		}
	}

	response.Image = productImages

	return response
}

func GetPaginationProducts(products []*entities.ProductEntity) []*TGetAllProductsResponse {
	var productFormatters []*TGetAllProductsResponse

	for _, product := range products {
		formattedProduct := GetAllProductsResponse(product)
		productFormatters = append(productFormatters, &formattedProduct)
	}

	return productFormatters
}

type CreateImageProductFormatter struct {
	Id  int    `json:"id"`
	Url string `json:"image"`
}

func CreateImageProductResponse(productImage *entities.ProductPhotosEntity) CreateImageProductFormatter {
	response := CreateImageProductFormatter{}
	response.Id = productImage.ID
	response.Url = productImage.ImageURL
	return response
}
