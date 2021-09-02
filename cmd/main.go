package main

import (
	"github.com/Shv-sergey70/todo-app-go"
	"github.com/Shv-sergey70/todo-app-go/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)
	server := new(todo.Server)

	if err := server.Run("8083", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server %s", err.Error())
	}
}
