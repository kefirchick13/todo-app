package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/kefirchick13/todo-app"
	"github.com/kefirchick13/todo-app/pkg/repository"
)

const (
	salt      = "sl;adk;sald;;nvvyeru"
	timeTTl   = 12 * time.Hour
	signInKey = "orfdsjl43785652312089"
)

type TokenClaims struct {
	jwt.StandardClaims
	UserId int
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = GeneratePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(UserName string, Password string) (string, error) {
	// get user from db
	user, err := s.repo.GetUser(UserName, GeneratePasswordHash(Password))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, TokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(timeTTl).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})
	return token.SignedString([]byte(signInKey))

}

func (s *AuthService) ParseToken(accessToken string) (int, error){
	token, err := jwt.ParseWithClaims(accessToken, &TokenClaims{}, func(token * jwt.Token) (interface{}, error){
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok{
			return nil, errors.New("invalid signig method")
		}
		return []byte(signInKey), nil
	})
	if err != nil{
		return 0, err
	}
	claims, ok := token.Claims.(*TokenClaims)
	if !ok{
		return 0, errors.New("token claims are not of type *TokenClaims")
	}
	return claims.UserId, err
}


func GeneratePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
