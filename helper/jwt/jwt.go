package jwt

import (
	"github.com/golang-jwt/jwt"
	"time"
)

type IJwt interface {
	GenerateJWT(userID int, email, role string) (string, error)
	ValidateToken(tokenString string) (*jwt.Token, error)
}

type JWT struct {
	Secret string
}

func NewJWT(secret string) IJwt {
	return &JWT{
		Secret: secret,
	}
}

func (j *JWT) GenerateJWT(userID int, email, role string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"email":   email,
		"role":    role,
		"exp":     time.Now().Add(time.Hour * 12).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := token.SignedString([]byte(j.Secret))
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func (j *JWT) ValidateToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(j.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}
