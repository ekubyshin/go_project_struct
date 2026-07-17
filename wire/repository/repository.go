package repository

import "fmt"

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
