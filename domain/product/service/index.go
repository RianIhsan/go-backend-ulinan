package service

import (
	"errors"
	"math"
	"ulinan/domain/category"
	"ulinan/domain/product"
	"ulinan/domain/product/dto"
	"ulinan/entities"
)

type ProductService struct {
	repo            product.ProductRepositoryInterface
	categoryService category.CategoryServiceInterface
}

func NewProductService(repo product.ProductRepositoryInterface, categoryService category.CategoryServiceInterface) product.ProductServiceInterface {
	return &ProductService{repo, categoryService}
}

func (s *ProductService) CreateProduct(request *dto.TCreateProductRequest) (*entities.ProductEntity, error) {
	_, err := s.categoryService.GetCategoryById(request.CategoryID)
	if err != nil {
		return nil, errors.New("category not found")
	}

	productData := &entities.ProductEntity{
		Name:        request.Name,
		Price:       request.Price,
		Description: request.Description,
		Address:     request.Address,
		CategoryID:  request.CategoryID,
	}

	createdProduct, err := s.repo.InsertProduct(productData)
	if err != nil {
		return nil, errors.New("failed to create product")
	}
	return createdProduct, nil
}

func (s *ProductService) GetAllProducts(page, perPage int) ([]*entities.ProductEntity, int64, error) {
	products, err := s.repo.FindProducts(page, perPage)
	if err != nil {
		return nil, 0, err
	}
	totalProducts, err := s.repo.CountTotalProducts()
	if err != nil {
		return nil, 0, err
	}
	return products, totalProducts, nil
}

func (s *ProductService) CalculatePaginationValues(page int, totalItems int, perPage int) (int, int) {
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

func (s *ProductService) GetNextPage(currentPage, totalPages int) int {
	if currentPage < totalPages {
		return currentPage + 1
	}
	return totalPages
}

func (s *ProductService) GetPrevPage(currentPage int) int {
	if currentPage > 1 {
		return currentPage - 1
	}
	return 1
}

func (s *ProductService) GetProductByName(page, perPage int, name string) ([]*entities.ProductEntity, int64, error) {
	products, err := s.repo.FindProductByName(page, perPage, name)
	if err != nil {
		return nil, 0, err
	}
	totalProducts, err := s.repo.CountProductByName(name)
	if err != nil {
		return nil, 0, err
	}
	return products, totalProducts, nil
}

func (s *ProductService) GetProductById(id int) (*entities.ProductEntity, error) {
	product, err := s.repo.FindProductById(id)
	if err != nil {
		return nil, errors.New("product not found")
	}
	return product, nil
}

func (s *ProductService) CreateProductImage(request dto.CreateProductImage) (*entities.ProductPhotosEntity, error) {
	product, err := s.GetProductById(request.ProductID)
	if err != nil {
		return nil, errors.New("product not found")
	}
	value := &entities.ProductPhotosEntity{
		ProductID: product.ID,
		ImageURL:  request.Image,
	}

	images, err := s.repo.CreateImageProduct(value)
	if err != nil {
		return nil, errors.New("failed to create product image")
	}

	return images, nil
}

func (s *ProductService) UpdateProduct(id int, request *dto.TUpdateProductRequest) error {
	product, err := s.repo.FindProductById(id)
	if err != nil {
		return errors.New("product not found")
	}

	if request.Name != "" {
		product.Name = request.Name
	}
	if request.Price != 0 {
		product.Price = request.Price
	}
	if request.Description != "" {
		product.Description = request.Description
	}
	if request.Address != "" {
		product.Address = request.Address
	}
	if request.CategoryID != 0 {
		_, err := s.categoryService.GetCategoryById(request.CategoryID)
		if err != nil {
			return errors.New("category not found")
		}
		product.CategoryID = request.CategoryID
	}

	_, err = s.repo.UpdateProduct(product.ID, product)
	if err != nil {
		return errors.New("failed to update product")
	}
	return nil
}

func (s *ProductService) DeleteProduct(id int) error {
	product, err := s.repo.FindProductById(id)
	if err != nil {
		return errors.New("product not found")
	}
	err = s.repo.DeleteProduct(product.ID)
	if err != nil {
		return errors.New("failed to delete product")
	}

	return nil
}

func (s *ProductService) DeleteProductImage(productId, ImageId int) error {
	product, err := s.repo.FindProductById(productId)
	if err != nil {
		return errors.New("product not found")
	}
	found := false
	for _, productImage := range product.ProductPhotos {
		if productImage.ID == ImageId {
			found = true
			break
		}
	}
	if !found {
		return errors.New("image not found")
	}
	if err := s.repo.DeleteProductImage(product.ID, ImageId); err != nil {
		return errors.New("failed to delete product image")
	}

	return nil
}

func (s *ProductService) GetRandomProducts(count int) ([]*entities.ProductEntity, error) {
	products, err := s.repo.GetRandomProducts(count)
	if err != nil {
		return nil, errors.New("failed to get random products")
	}
	return products, nil
}

func (s *ProductService) GetProductByCategory(page, perPage, categoryId int) ([]*entities.ProductEntity, int64, error) {
	products, err := s.repo.GetProductsByCategoryID(page, perPage, categoryId)
	if err != nil {
		return nil, 0, err
	}
	totalProducts, err := s.repo.CountTotalProductsByCategoryID(categoryId)
	if err != nil {
		return nil, 0, err
	}
	return products, totalProducts, nil
}

func (s *ProductService) GetProductsByCategoryAndNameProduct(page, perPage, categoryID int, search string) ([]*entities.ProductEntity, int64, error) {
	products, err := s.repo.GetProductsByCategoryAndNameProduct(page, perPage, categoryID, search)
	if err != nil {
		return nil, 0, err
	}

	totalProducts, err := s.repo.CountProductByCategoryAndNameProduct(categoryID, search)
	if err != nil {
		return nil, 0, err
	}

	return products, totalProducts, nil
}
