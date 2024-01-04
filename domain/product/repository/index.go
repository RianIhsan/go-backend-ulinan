package repository

import (
	"gorm.io/gorm"
	"time"
	"ulinan/domain/product"
	"ulinan/entities"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) product.ProductRepositoryInterface {
	return &ProductRepository{db}
}

func (r *ProductRepository) InsertProduct(productData *entities.ProductEntity) (*entities.ProductEntity, error) {
	if err := r.db.Create(productData).Error; err != nil {
		return nil, err
	}
	return productData, nil
}

func (r *ProductRepository) FindProducts(page, perPage int) ([]*entities.ProductEntity, error) {
	var products []*entities.ProductEntity
	offset := (page - 1) * perPage
	err := r.db.Offset(offset).Limit(perPage).Preload("Category").Preload("ProductPhotos").Where("deleted_at IS NULL").Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (r *ProductRepository) CountTotalProducts() (int64, error) {
	var count int64
	err := r.db.Model(&entities.ProductEntity{}).Preload("Category").Preload("ProductPhotos").Where("deleted_at IS NULL").Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *ProductRepository) FindProductById(id int) (*entities.ProductEntity, error) {
	var product *entities.ProductEntity
	err := r.db.Preload("Category").Preload("ProductPhotos").Where("id = ? AND deleted_at IS NULL", id).First(&product).Error
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (r *ProductRepository) FindProductByName(page, perPage int, name string) ([]*entities.ProductEntity, error) {
	var products []*entities.ProductEntity
	offset := (page - 1) * perPage
	query := r.db.Offset(offset).Limit(perPage).Preload("Category").Preload("ProductPhotos").Where("deleted_at IS NULL")
	if name != "" {
		query = query.Where("name LIKE ?", "%"+name+"%")
	}
	err := query.Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (r *ProductRepository) CountProductByName(name string) (int64, error) {
	var count int64
	query := r.db.Model(&entities.ProductEntity{}).Preload("Category").Preload("ProductPhotos").Where("deleted_at IS NULL")
	if name != "" {
		query = query.Where("name LIKE ?", "%"+name+"%")
	}
	err := query.Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *ProductRepository) CreateImageProduct(productImage *entities.ProductPhotosEntity) (*entities.ProductPhotosEntity, error) {
	if err := r.db.Create(productImage).Error; err != nil {
		return nil, err
	}
	return productImage, nil
}

func (r *ProductRepository) UpdateProduct(id int, updatedProduct *entities.ProductEntity) (*entities.ProductEntity, error) {
	var product *entities.ProductEntity
	err := r.db.Where("id = ? AND deleted_at IS NULL", id).First(&product).Error
	if err != nil {
		return nil, err
	}
	if err := r.db.Model(&product).Updates(updatedProduct).Error; err != nil {
		return nil, err
	}

	return product, err
}

func (r *ProductRepository) DeleteProduct(id int) error {
	product := &entities.ProductEntity{}
	if err := r.db.First(product, id).Error; err != nil {
		return err
	}
	if err := r.db.Model(product).Update("deleted_at", time.Now()).Error; err != nil {
		return err
	}
	return nil
}

func (r *ProductRepository) DeleteProductImage(productId, ImageId int) error {
	tx := r.db.Begin()
	if err := tx.Model(&entities.ProductPhotosEntity{}).Where("id = ? AND product_id = ?", ImageId, productId).Update("deleted_at", time.Now()).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Model(&entities.ProductEntity{}).Where("id = ?", productId).Update("updated_at", time.Now()).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
