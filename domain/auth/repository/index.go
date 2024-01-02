package repository

import (
	"gorm.io/gorm"
	"ulinan/domain/auth"
	"ulinan/entities"
)

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) auth.AuthRepositoryInterface {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) InsertUser(newUser *entities.UserEntity) (*entities.UserEntity, error) {
	if err := r.db.Create(newUser).Error; err != nil {
		return nil, err
	}
	return newUser, nil
}

func (r *AuthRepository) FindUserByEmailOrUsername(identifier string) (*entities.UserEntity, error) {
	var user *entities.UserEntity
	if err := r.db.Table("users").
		Where("(email = ? OR username = ?) AND deleted_at IS NULL", identifier, identifier).
		First(&user).
		Error; err != nil {
		return nil, err
	}
	return user, nil
}
