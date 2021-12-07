package auth

import (
	"errors"
	"github.com/golang-jwt/jwt"
)

type Service interface {
	GenerateToken(userID int) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type JwtService struct {
}

func NewService() *JwtService {
	return &JwtService{}
}

var SecretKey = []byte("r3z4_s3cr3T_k3Y")

func (s *JwtService) GenerateToken(userID int) (string, error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err := token.SignedString(SecretKey)

	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}

func (s *JwtService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("Invalid token")
		}

		return []byte(SecretKey), nil
	})

	if err != nil {
		return token, err
	}

	return token, nil
}
