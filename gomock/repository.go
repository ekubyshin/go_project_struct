package gomockexample

import "fmt"

// UserRepository — интерфейс, от которого зависит сервис.
// В тестах подставляем mock вместо реальной реализации.
type UserRepository interface {
	FindByID(id int) (string, error)
}

// InMemoryUserRepository — простая реализация для production-кода.
type InMemoryUserRepository struct {
	users map[int]string
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users: map[int]string{
			1: "Алиса",
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
