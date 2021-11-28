package auth

import (
	"bwastartup/helper"
	"errors"
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

type Service interface {
	GenerateToken(id int) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtService struct {
}

func NewServiceJwt() *jwtService {
	return &jwtService{}
}

var SECRET_KEY = []byte(helper.EnvVariable("SECRET_KEY"))

func (s *jwtService) GenerateToken(id int) (string, error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = id

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	jwtToken, err := token.SignedString(SECRET_KEY)

	if err != nil {
		fmt.Println(err.Error())
		return jwtToken, err
	}
	return jwtToken, err
}

func (s *jwtService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {

		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("Invalid Token")
		}
		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		return token, err
	}
	return token, nil

}
