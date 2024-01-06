package product

import (
	"github.com/gofiber/fiber/v2"
	"ulinan/domain/product/dto"
	"ulinan/entities"
)

type ProductRepositoryInterface interface {
	InsertProduct(productData *entities.ProductEntity) (*entities.ProductEntity, error)
	FindProducts(page, perPage int) ([]*entities.ProductEntity, error)
	CountTotalProducts() (int64, error)
	FindProductById(id int) (*entities.ProductEntity, error)
	FindProductByName(page, perPage int, name string) ([]*entities.ProductEntity, error)
	CountProductByName(name string) (int64, error)
	CreateImageProduct(productImage *entities.ProductPhotosEntity) (*entities.ProductPhotosEntity, error)
	UpdateProduct(id int, updatedProduct *entities.ProductEntity) (*entities.ProductEntity, error)
	DeleteProduct(id int) error
	DeleteProductImage(productId, ImageId int) error
	GetRandomProducts(count int) ([]*entities.ProductEntity, error)
	GetProductsByCategoryID(page, perPage, categoryID int) ([]*entities.ProductEntity, error)
	CountTotalProductsByCategoryID(categoryID int) (int64, error)
	GetProductsByCategoryAndNameProduct(page, perPage, categoryID int, search string) ([]*entities.ProductEntity, error)
	CountProductByCategoryAndNameProduct(categoryID int, search string) (int64, error)
}

type ProductServiceInterface interface {
	CreateProduct(request *dto.TCreateProductRequest) (*entities.ProductEntity, error)
	GetAllProducts(page, perPage int) ([]*entities.ProductEntity, int64, error)
	CalculatePaginationValues(page int, totalItems int, perPage int) (int, int)
	GetNextPage(currentPage, totalPages int) int
	GetPrevPage(currentPage int) int
	GetProductByName(page, perPage int, name string) ([]*entities.ProductEntity, int64, error)
	GetProductById(id int) (*entities.ProductEntity, error)
	CreateProductImage(request dto.CreateProductImage) (*entities.ProductPhotosEntity, error)
	UpdateProduct(id int, req *dto.TUpdateProductRequest) error
	DeleteProduct(id int) error
	DeleteProductImage(productId, ImageId int) error
	GetRandomProducts(count int) ([]*entities.ProductEntity, error)
	GetProductByCategory(page, perPage, categoryId int) ([]*entities.ProductEntity, int64, error)
	GetProductsByCategoryAndNameProduct(page, perPage, categoryID int, search string) ([]*entities.ProductEntity, int64, error)
}

type ProductHandlerInterface interface {
	CreateProduct(c *fiber.Ctx) error
	GetProductById(c *fiber.Ctx) error
	GetAllProducts(c *fiber.Ctx) error
	GetRandomProducts(c *fiber.Ctx) error
	CreateProductImage(c *fiber.Ctx) error
	UpdateProduct(c *fiber.Ctx) error
	DeleteProduct(c *fiber.Ctx) error
	DeleteProductImage(c *fiber.Ctx) error
}
