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
	if err := r.db.Preload("Gender"). // Preload gender
						Where("id = ? AND deleted_at IS NULL", id).
						First(&user).
						Error; err != nil {
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

func (r *UserRepository) UpdateUserAvatar(userID int, avatarPath string) error {
	user := &entities.UserEntity{ID: userID}

	result := r.db.Model(user).Updates(map[string]interface{}{"avatar": avatarPath})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *UserRepository) UpdateUser(userUpdate *entities.UserEntity) error {
	result := r.db.Preload("Gender").
		Model(userUpdate).
		Updates(userUpdate)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *UserRepository) UpdateGender(user *entities.UserEntity) (*entities.UserEntity, error) {
	result := r.db.Preload("Gender").
		Where("id = ?", user.ID).
		Find(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (r *UserRepository) UpdatePassword(userUpdate *entities.UserEntity) error {
	result := r.db.Model(userUpdate).Updates(map[string]interface{}{"password": userUpdate.Password})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
