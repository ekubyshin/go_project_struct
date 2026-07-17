// Пример 4: внедрение зависимостей через паттерн Builder.
//
// Builder пошагово собирает объект и проверяет, что все обязательные
// зависимости переданы до вызова Build().
package main

import (
	"errors"
	"fmt"
	"strings"
)

type UserRepository interface {
	FindByID(id int) (string, error)
}

type InMemoryUserRepository struct {
	users map[int]string
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users: map[int]string{
			1: "Алиса",
			2: "Боб",
		},
	}
}

func (r *InMemoryUserRepository) FindByID(id int) (string, error) {
	name, ok := r.users[id]
	if !ok {
		return "", fmt.Errorf("пользователь %d не найден", id)
	}
	return name, nil
}

type UserService struct {
	repo UserRepository
}

func (s *UserService) Greet(id int) (string, error) {
	name, err := s.repo.FindByID(id)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Привет, %s!", name), nil
}

// UserServiceBuilder накапливает зависимости и создаёт сервис.
type UserServiceBuilder struct {
	repo UserRepository
}

func NewUserServiceBuilder() *UserServiceBuilder {
	return &UserServiceBuilder{}
}

func (b *UserServiceBuilder) WithRepository(repo UserRepository) *UserServiceBuilder {
	b.repo = repo
	return b
}

func (b *UserServiceBuilder) Build() (*UserService, error) {
	if b.repo == nil {
		return nil, errors.New("обязательная зависимость UserRepository не задана")
	}
	return &UserService{repo: b.repo}, nil
}

func main() {
	repo := NewInMemoryUserRepository()

	// Цепочка вызовов: builder принимает зависимости и валидирует их в Build().
	service, err := NewUserServiceBuilder().
		WithRepository(repo).
		Build()
	if err != nil {
		panic(err)
	}

	msg, err := service.Greet(1)
	if err != nil {
		panic(err)
	}

	fmt.Println(strings.TrimSpace(msg))
}
