package category

import (
	"github.com/gofiber/fiber/v2"
	"ulinan/domain/category/dto"
	"ulinan/entities"
)

type CategoryRepositoryInterface interface {
	Insert(category *entities.CategoryEntity) (*entities.CategoryEntity, error)
	FindCategories(page, perPage int) ([]*entities.CategoryEntity, error)
	CountTotalCategories() (int64, error)
	FindCategoryById(id int) (*entities.CategoryEntity, error)
	UpdateCategory(id int, updatedCategory *entities.CategoryEntity) (*entities.CategoryEntity, error)
	DeleteCategory(id int) error
	FindByName(page, perPage int, name string) ([]*entities.CategoryEntity, error)
	CountCategoryByName(name string) (int64, error)
}

type CategoryServiceInterface interface {
	CreateCategory(request *dto.TCreateCategoryRequest) (*entities.CategoryEntity, error)
	GetAllCategories(page, perPage int) ([]*entities.CategoryEntity, int64, error)
	CalculatePaginationValues(page int, totalItems int, perPage int) (int, int)
	GetNextPage(currentPage, totalPages int) int
	GetPrevPage(currentPage int) int
	UpdateCategory(id int, req *dto.TUpdateCategoryRequest) error
	DeleteCategory(id int) error
	GetCategoryByName(page, perPage int, name string) ([]*entities.CategoryEntity, int64, error)
	GetCategoryById(id int) (*entities.CategoryEntity, error)
}

type CategoryHandlerInterface interface {
	CreateCategory(c *fiber.Ctx) error
	GetCategoryById(c *fiber.Ctx) error
	GetAllCategories(c *fiber.Ctx) error
	UpdateCategory(c *fiber.Ctx) error
	DeleteCategory(c *fiber.Ctx) error
}
