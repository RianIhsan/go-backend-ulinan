package service

import (
	"errors"
	"time"
	"ulinan/domain/user"
	"ulinan/domain/user/dto"
	"ulinan/entities"
	"ulinan/helper/hashing"
)

type UserService struct {
	repo    user.UserRepositoryInterface
	hashing hashing.HashInterface
}

func NewUserService(repo user.UserRepositoryInterface, hashing hashing.HashInterface) user.UserServiceInterface {
	return &UserService{repo: repo, hashing: hashing}
}

func (s *UserService) GetEmail(email string) (*entities.UserEntity, error) {
	result, err := s.repo.FindEmail(email)
	if err != nil {
		return nil, errors.New("your email has been already")
	}
	return result, nil
}

func (s *UserService) GetId(id int) (*entities.UserEntity, error) {
	result, err := s.repo.FindId(id)
	if err != nil {
		return nil, errors.New("id not found")
	}
	return result, nil
}

func (s *UserService) GetUsername(username string) (*entities.UserEntity, error) {
	result, err := s.repo.FindUsername(username)
	if err != nil {
		return nil, errors.New("username not found")
	}
	return result, nil
}

func (s *UserService) UpdateAvatar(userID int, request *dto.TUpdateAvatarRequest) (*entities.UserEntity, error) {
	user, err := s.repo.FindId(userID)
	if err != nil {
		return nil, errors.New("failed to get user")
	}

	if user == nil {
		return nil, errors.New("user not found")
	}

	userUpdateAvatar := &entities.UserEntity{
		Avatar: request.Avatar,
	}

	err = s.repo.UpdateUserAvatar(userID, userUpdateAvatar.Avatar)
	if err != nil {
		return nil, errors.New("failed to update user avatar")
	}

	user.Avatar = userUpdateAvatar.Avatar

	return user, nil
}

func (s *UserService) UpdatePassword(userID int, request *dto.TUpdatePasswordRequest) error {
	userToUpdate, err := s.repo.FindId(userID)
	if err != nil {
		return errors.New("failed to get user")
	}

	if userToUpdate == nil {
		return errors.New("user not found")
	}

	isOldPasswordValid, err := s.hashing.ComparePassword(userToUpdate.Password, request.OldPassword)
	if err != nil || !isOldPasswordValid {
		return errors.New("incorrect old password")
	}

	if request.NewPassword != request.ConfirmPassword {
		return errors.New("password tidak tidak sesuai")
	}

	newHashedPassword, err := s.hashing.GenerateHash(request.NewPassword)
	if err != nil {
		return errors.New("failed to generate new password hash")
	}

	userToUpdate.Password = newHashedPassword

	err = s.repo.UpdatePassword(userToUpdate)
	if err != nil {
		return errors.New("failed to update password")
	}

	return nil
}

func (s *UserService) UpdateUser(userID int, request *dto.TUpdateUserRequest) (*entities.UserEntity, error) {
	userToUpdate, err := s.repo.FindId(userID)
	if err != nil {
		return nil, errors.New("failed to get user")
	}

	if userToUpdate == nil {
		return nil, errors.New("user not found")
	}

	if request.Fullname != nil {
		userToUpdate.Fullname = *request.Fullname
	}
	if request.Username != nil {
		userToUpdate.Username = *request.Username
	}
	if request.Email != nil {
		userToUpdate.Email = *request.Email
	}
	if request.Phone != nil {
		userToUpdate.Phone = request.Phone
	}
	if request.Birthdate != nil {
		birthdate, err := time.Parse("2006-01-02", *request.Birthdate)
		if err != nil {
			return nil, errors.New("invalid birthdate format")
		}
		userToUpdate.Birthdate = &birthdate
	}
	if request.Address != nil {
		userToUpdate.Address = request.Address
	}
	if request.GenderID != nil {
		if userToUpdate.Gender == nil {
			userToUpdate.Gender = &entities.GenderEntity{}
		}
		userToUpdate.Gender.ID = *request.GenderID
	}

	err = s.repo.UpdateUser(userToUpdate)
	if err != nil {
		return nil, errors.New("failed to update user")
	}

	return userToUpdate, nil
}

func (s *UserService) UpdateUserGender(userID int, genderID int) error {
	userToUpdate, err := s.repo.FindId(userID)
	if err != nil {
		return errors.New("failed to get user")
	}

	if userToUpdate == nil {
		return errors.New("user not found")
	}

	err = s.repo.UpdateUserGenderWithTransaction(userToUpdate, genderID)

	if err != nil {
		return err
	}

	return nil
}
