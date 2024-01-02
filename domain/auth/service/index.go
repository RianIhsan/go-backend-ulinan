package service

import (
	"errors"
	"ulinan/domain/auth"
	"ulinan/domain/auth/dto"
	"ulinan/domain/user"
	"ulinan/entities"
	"ulinan/helper/hashing"
	"ulinan/helper/jwt"
)

type AuthService struct {
	repo        auth.AuthRepositoryInterface
	userService user.UserServiceInterface
	hashing     hashing.HashInterface
	jwt         jwt.IJwt
}

func NewAuthService(repo auth.AuthRepositoryInterface, userService user.UserServiceInterface, hashing hashing.HashInterface, jwt jwt.IJwt) auth.AuthServiceInterface {
	return &AuthService{repo: repo, userService: userService, hashing: hashing, jwt: jwt}
}

func (s *AuthService) Register(payload *dto.TRegisterRequest) (*entities.UserEntity, error) {
	isExistEmail, _ := s.userService.GetEmail(payload.Email)
	if isExistEmail != nil {
		return nil, errors.New("email already exists")
	}

	isExistUsername, _ := s.userService.GetUsername(payload.Username)
	if isExistUsername != nil {
		return nil, errors.New("username already exists")
	}

	if payload.Password != payload.PasswordConfirm {
		return nil, errors.New("password does not match")
	}

	hashPassword, err := s.hashing.GenerateHash(payload.Password)
	if err != nil {
		return nil, err
	}

	newUser := &entities.UserEntity{
		Fullname: payload.Fullname,
		Username: payload.Username,
		Avatar:   "https://res.cloudinary.com/dyominih0/image/upload/v1702051015/my-sample-avatar/voaa01wefhnziwzqwn1m.webp",
		Email:    payload.Email,
		Password: hashPassword,
		Role:     "user",
	}

	user, err := s.repo.InsertUser(newUser)
	if err != nil {
		return nil, errors.New("failed create account")
	}

	return user, nil
}

func (s *AuthService) Login(payload *dto.TLoginRequest) (*entities.UserEntity, string, error) {
	user, err := s.repo.FindUserByEmailOrUsername(payload.Identifier)
	if err != nil {
		return nil, "", errors.New("user not found")
	}

	isValidPassword, err := s.hashing.ComparePassword(user.Password, payload.Password)
	if err != nil || !isValidPassword {
		return nil, "", errors.New("incorrect password")
	}

	accessSecret, err := s.jwt.GenerateJWT(user.ID, user.Email, user.Role)
	if err != nil {
		return nil, "", err
	}

	return user, accessSecret, nil
}
