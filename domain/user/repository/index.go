package repository

import (
	"errors"
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

func (r *UserRepository) UpdateGender(user *entities.UserEntity) error {
	if err := r.db.Save(user).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) UpdatePassword(userUpdate *entities.UserEntity) error {
	result := r.db.Model(userUpdate).Updates(map[string]interface{}{"password": userUpdate.Password})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *UserRepository) FindGenderByID(genderID int) (*entities.GenderEntity, error) {
	var gender entities.GenderEntity
	if err := r.db.First(&gender, genderID).Error; err != nil {
		return nil, err
	}
	return &gender, nil
}

func (r *UserRepository) UpdateUserWithTransaction(user *entities.UserEntity) (*entities.UserEntity, error) {
	tx := r.db.Begin()

	if err := tx.Save(user).Error; err != nil {
		tx.Rollback()
		return nil, errors.New("gagal menyimpan user: " + err.Error())
	}

	return user, tx.Commit().Error
}

func (r *UserRepository) UpdateUserGenderWithTransaction(user *entities.UserEntity, genderID int) error {
	tx := r.db.Begin()

	if user.GenderID != nil {
		// Delete existing gender_id
		if err := tx.Model(user).UpdateColumn("gender_id", nil).Error; err != nil {
			tx.Rollback()
			return errors.New("failed to delete existing gender_id: " + err.Error())
		}
	}

	// Set the new gender_id
	if err := tx.Model(user).Update("gender_id", genderID).Error; err != nil {
		tx.Rollback()
		return errors.New("failed to set new gender_id: " + err.Error())
	}

	return tx.Commit().Error
}
