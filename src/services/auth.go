package services

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"log"
	"time"

	"example.com/hello/src/models"
	"example.com/hello/src/repository"
	"github.com/dgrijalva/jwt-go"
)

const (
	salt = "jkmng7)mdm,^Damn_!"
	sign = "fj9fmckll"
)

type AuthService struct {
	repo            repository.Authorization
	rabbitTransport repository.RabbitTransport
}

type TokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

func NewAuthService(repo *repository.Repository) *AuthService {
	return &AuthService{
		repo:            repo.Authorization,
		rabbitTransport: repo.RabbitTransport,
	}
}

func (as *AuthService) ParseToken(accesstoken string) (int, error) {
	token, err := jwt.ParseWithClaims(accesstoken, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(sign), nil

	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*TokenClaims)

	if !ok {
		return 0, errors.New("token claims corrupted")
	}

	return claims.UserId, nil

}

func (as *AuthService) CreateUser(user models.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return as.repo.CreateUser(user)
}

func (as *AuthService) LoginUser(user models.LoginUserStruct) (string, error) {
	dbUser, error := as.repo.GetUserFromDb(user.Username, generatePasswordHash(user.Password))

	if error != nil {
		return "", error
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		}, dbUser.Id,
	})

	tokenByte, error := token.SignedString([]byte(sign))

	if error != nil {
		log.Printf("token creation failed")
	}

	go as.rabbitTransport.PushPackage(tokenByte)

	if error != nil {
		log.Printf("token sending failed")
	}

	return tokenByte, error
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
