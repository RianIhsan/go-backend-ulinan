package service

import (
	"errors"
	"math"
	"ulinan/domain/category"
	"ulinan/domain/category/dto"
	"ulinan/entities"
)

type CategoryService struct {
	repo category.CategoryRepositoryInterface
}

func NewCategoryService(repo category.CategoryRepositoryInterface) category.CategoryServiceInterface {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) CreateCategory(request *dto.TCreateCategoryRequest) (*entities.CategoryEntity, error) {
	result := &entities.CategoryEntity{
		Name:        request.Name,
		Description: request.Description,
		Image:       request.Image,
	}

	createdCategory, err := s.repo.Insert(result)
	if err != nil {
		return nil, err
	}

	return createdCategory, nil
}

func (s *CategoryService) GetAllCategories(page, perPage int) ([]*entities.CategoryEntity, int64, error) {
	categoryList, err := s.repo.FindCategories(page, perPage)
	if err != nil {
		return nil, 0, err
	}
	totalCategories, err := s.repo.CountTotalCategories()
	if err != nil {
		return nil, 0, err
	}
	return categoryList, totalCategories, nil
}

func (s *CategoryService) CalculatePaginationValues(page int, totalItems int, perPage int) (int, int) {
	pageInt := page
	if pageInt <= 0 {
		pageInt = 1
	}

	totalPages := int(math.Ceil(float64(totalItems) / float64(perPage)))

	if pageInt > totalPages {
		pageInt = totalPages
	}

	return pageInt, totalPages
}

func (s *CategoryService) GetNextPage(currentPage, totalPages int) int {
	if currentPage < totalPages {
		return currentPage + 1
	}
	return totalPages
}

func (s *CategoryService) GetPrevPage(currentPage int) int {
	if currentPage > 1 {
		return currentPage - 1
	}
	return 1
}

func (s *CategoryService) UpdateCategory(id int, req *dto.TUpdateCategoryRequest) error {
	category, err := s.repo.FindCategoryById(id)
	if err != nil {
		return errors.New("category not found")
	}

	if req.Name != "" {
		category.Name = req.Name
	}
	if req.Description != "" {
		category.Description = req.Description
	}
	if req.Image != "" {
		category.Image = req.Image
	}
	_, err = s.repo.UpdateCategory(category.ID, category)
	if err != nil {
		return err
	}
	return nil
}

func (s *CategoryService) DeleteCategory(id int) error {
	category, err := s.repo.FindCategoryById(id)
	if err != nil {
		return errors.New("category not found")
	}

	err = s.repo.DeleteCategory(category.ID)
	if err != nil {
		return err
	}
	return nil
}

func (s *CategoryService) GetCategoryByName(page, perPage int, name string) ([]*entities.CategoryEntity, int64, error) {
	categoryList, err := s.repo.FindByName(page, perPage, name)
	if err != nil {
		return nil, 0, err
	}
	totalCategories, err := s.repo.CountCategoryByName(name)
	if err != nil {
		return nil, 0, err
	}
	return categoryList, totalCategories, nil
}

func (s *CategoryService) GetCategoryById(id int) (*entities.CategoryEntity, error) {
	category, err := s.repo.FindCategoryById(id)
	if err != nil {
		return nil, err
	}
	return category, nil
}
