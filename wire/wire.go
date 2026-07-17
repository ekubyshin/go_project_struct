//go:build wireinject
// +build wireinject

// Пример 2: автоматическая сборка зависимостей через Google Wire.
//
// Wire — это code generator, а не runtime-фреймворк.
// На этапе компиляции он генерирует обычный Go-код с явными вызовами конструкторов.
package main

import (
	"github.com/example/dependency/wire/repository"
	"github.com/example/dependency/wire/service"
	"github.com/google/wire"
)

// wire.Build описывает, какие конструкторы использовать.
// Wire сам разберётся, что UserService нужен UserRepository.
func InitializeUserService() (*service.UserService, error) {
	wire.Build(
		repository.NewInMemoryUserRepository,
		wire.Bind(new(repository.UserRepository), new(*repository.InMemoryUserRepository)),
		service.NewUserService,
	)
	return nil, nil
}
