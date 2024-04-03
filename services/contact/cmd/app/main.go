package main

import (
	"fmt"
	postgres "murat/clean_architecture/pkg/store/postres"
)

func main() {
	store, err := postgres.New(postgres.Settings{
		Host:     "localhost",
		Port:     5432,
		Database: "test",
		User:     "user",
		Password: "password",
		SSLMode:  false,
	})

	if err != nil {
		panic(err)
	}

	defer store.Pool.Close()

	store.Pool.Stat()

	fmt.Println("Hello World")
}
