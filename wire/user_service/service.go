package service

import (
	"fmt"

	"github.com/example/dependency/wire/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Greet(id int) (string, error) {
	name, err := s.repo.FindByID(id)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Привет, %s!", name), nil
}
