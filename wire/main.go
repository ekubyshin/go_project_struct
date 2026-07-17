package main

import (
	"fmt"
	"strings"
)

func main() {
	// Сборка графа зависимостей — одна строка.
	// Реализация сгенерирована в wire_gen.go командой: go generate ./wire
	service, err := InitializeUserService()
	if err != nil {
		panic(err)
	}

	msg, err := service.Greet(1)
	if err != nil {
		panic(err)
	}

	fmt.Println(strings.TrimSpace(msg))
}

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
