package service

import (
	"time"

	"github.com/gbrlsnchs/jwt/v3"
	config "github.com/synergydesigns/stylesblitz-server/shared/config"
	models "github.com/synergydesigns/stylesblitz-server/shared/models"
	"golang.org/x/crypto/bcrypt"
)

type JWTService struct {
	hash           jwt.Algorithm
	passwordSecret string
}

type JWT interface {
	GenerateAuthToken(user models.User) (string, error)
	DecodeToken(token string) (models.User, error)
	HashPassword(password string) (string, error)
	CheckPasswordHash(password, hash string) bool
}

type JWTPayload struct {
	jwt.Payload
	User models.User
}

func NewJWT(conf *config.Config) *JWTService {
	var hash = jwt.NewHS256([]byte(conf.AuthenticationSecret))

	jwt := JWTService{
		hash:           hash,
		passwordSecret: conf.PasswordSecret,
	}

	return &jwt
}

func (service *JWTService) GenerateAuthToken(user models.User) (string, error) {
	now := time.Now()
	tokenPayload := JWTPayload{
		Payload: jwt.Payload{
			Issuer:         "lookblitz",
			Audience:       jwt.Audience{"localhost:3001"},
			ExpirationTime: jwt.NumericDate(now.Add(24 * 30 * 12 * time.Hour)),
			NotBefore:      jwt.NumericDate(now.Add(30 * time.Minute)),
			IssuedAt:       jwt.NumericDate(now),
		},
		User: user,
	}

	token, err := jwt.Sign(tokenPayload, service.hash)
	if err != nil {
		return "", err
	}

	return string(token), nil
}

func (service *JWTService) DecodeToken(token string) (models.User, error) {
	var payload JWTPayload

	_, err := jwt.Verify([]byte(token), service.hash, &payload)

	if err != nil {
		return payload.User, err
	}

	return payload.User, nil
}

func (service *JWTService) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func (service *JWTService) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
