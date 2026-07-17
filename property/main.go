// Пример 3: внедрение зависимостей через свойство (property injection).
//
// Зависимость задаётся после создания объекта — через сеттер или публичное поле.
// В Go это встречается реже конструкторного DI: нет гарантии на этапе компиляции,
// что зависимость уже установлена.
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

// UserService создаётся пустым, репозиторий подставляется позже.
type UserService struct {
	repo UserRepository
}

func NewUserService() *UserService {
	return &UserService{}
}

// SetRepository — сеттер для property injection.
func (s *UserService) SetRepository(repo UserRepository) {
	s.repo = repo
}

func (s *UserService) Greet(id int) (string, error) {
	if s.repo == nil {
		return "", errors.New("репозиторий не задан: вызовите SetRepository")
	}

	name, err := s.repo.FindByID(id)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Привет, %s!", name), nil
}

func main() {
	repo := NewInMemoryUserRepository()

	// 1. Создаём объект
	service := NewUserService()

	// 2. Внедряем зависимость через свойство
	service.SetRepository(repo)

	msg, err := service.Greet(1)
	if err != nil {
		panic(err)
	}

	fmt.Println(strings.TrimSpace(msg))
}
