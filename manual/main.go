// Пример 1: внедрение зависимостей вручную, без фреймворков.
//
// Идея: зависимости передаются явно через конструкторы.
// main собирает граф объектов сам — это и есть "composition root".
package main

import (
	"fmt"
	"strings"
)

// --- Реализация ---

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

// --- Сервис зависит от интерфейса, а не от конкретной реализации ---



// --- Точка сборки приложения ---

func main() {
	// 1. Создаём зависимости снизу вверх
	repo := NewInMemoryUserRepository()
	service := NewUserService(repo)

	// 2. Используем готовый сервис
	msg, err := service.Greet(1)
	if err != nil {
		panic(err)
	}

	fmt.Println(strings.TrimSpace(msg))
}
