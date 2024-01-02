package repository

import (
	"gorm.io/gorm"
	"ulinan/domain/user"
	"ulinan/entities"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) user.UserRepositoryInterface {
	return &UserRepository{db: db}
}

func (r *UserRepository) FindEmail(email string) (*entities.UserEntity, error) {
	var user *entities.UserEntity
	if err := r.db.Table("users").
		Where("email = ? AND deleted_at IS NULL", email).
		First(&user).
		Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) FindId(id int) (*entities.UserEntity, error) {
	var user *entities.UserEntity
	if err := r.db.Where("id = ? AND deleted_at IS NULL", id).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) FindUsername(username string) (*entities.UserEntity, error) {
	var user *entities.UserEntity
	if err := r.db.Where("username = ? AND deleted_at IS NULL", username).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
