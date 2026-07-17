package gomockexample

import (
	"errors"
	"testing"

	"github.com/example/dependency/gomock/mocks"
	"go.uber.org/mock/gomock"
)

func TestUserService_Greet_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepository(ctrl)
	mockRepo.EXPECT().
		FindByID(1).
		Return("Алиса", nil)

	service := NewUserService(mockRepo)

	got, err := service.Greet(1)
	if err != nil {
		t.Fatalf("не ожидали ошибку: %v", err)
	}

	want := "Привет, Алиса!"
	if got != want {
		t.Fatalf("получили %q, хотели %q", got, want)
	}
}

func TestUserService_Greet_RepositoryError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepository(ctrl)
	mockRepo.EXPECT().
		FindByID(99).
		Return("", errors.New("пользователь не найден"))

	service := NewUserService(mockRepo)

	_, err := service.Greet(99)
	if err == nil {
		t.Fatal("ожидали ошибку от репозитория")
	}
}
