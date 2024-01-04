package repository

import (
	"gorm.io/gorm"
	"time"
	"ulinan/domain/category"
	"ulinan/entities"
)

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) category.CategoryRepositoryInterface {
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) Insert(category *entities.CategoryEntity) (*entities.CategoryEntity, error) {
	if err := r.db.Create(category).Error; err != nil {
		return nil, err
	}
	return category, nil
}

func (r *CategoryRepository) FindCategories(page, perPage int) ([]*entities.CategoryEntity, error) {
	var categories []*entities.CategoryEntity
	offset := (page - 1) * perPage
	err := r.db.Offset(offset).Limit(perPage).Where("deleted_at IS NULL").Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *CategoryRepository) CountTotalCategories() (int64, error) {
	var count int64
	err := r.db.Model(&entities.CategoryEntity{}).Where("deleted_at IS NULL").Count(&count).Error
	return count, err
}

func (r *CategoryRepository) FindCategoryById(id int) (*entities.CategoryEntity, error) {
	var category *entities.CategoryEntity
	err := r.db.Where("id = ? AND deleted_at IS NULL", id).First(&category).Error
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (r *CategoryRepository) UpdateCategory(id int, updatedCategory *entities.CategoryEntity) (*entities.CategoryEntity, error) {
	var category *entities.CategoryEntity
	err := r.db.Where("id = ? AND deleted_at IS NULL", id).First(&category).Error
	if err != nil {
		return nil, err
	}
	if err := r.db.Model(&category).Updates(updatedCategory).Error; err != nil {
		return nil, err
	}
	return category, nil
}

func (r *CategoryRepository) DeleteCategory(id int) error {
	category := &entities.CategoryEntity{}
	if err := r.db.First(category, id).Error; err != nil {
		return err
	}
	if err := r.db.Model(category).Update("deleted_at", time.Now()).Error; err != nil {
		return err
	}
	return nil
}

func (r *CategoryRepository) FindByName(page, perPage int, name string) ([]*entities.CategoryEntity, error) {
	var categories []*entities.CategoryEntity
	offset := (page - 1) * perPage
	query := r.db.Offset(offset).Limit(perPage).Where("deleted_at IS NULL")
	if name != "" {
		query = query.Where("name LIKE ?", "%"+name+"%")
	}
	err := query.Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *CategoryRepository) CountCategoryByName(name string) (int64, error) {
	var count int64
	query := r.db.Model(&entities.CategoryEntity{}).Where("deleted_at IS NULL")
	if name != "" {
		query = query.Where("name LIKE ?", "%"+name+"%")
	}
	err := query.Count(&count).Error
	return count, err
}
