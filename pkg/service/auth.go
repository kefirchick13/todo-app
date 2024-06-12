package service

import (
	"crypto/sha1"
	"fmt"

	"github.com/kefirchick13/todo-app"
	"github.com/kefirchick13/todo-app/pkg/repository"
)

const salt = "sl;adk;sald;;nvvyeru"

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

func GeneratePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
