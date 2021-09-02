package main

import (
	"github.com/Shv-sergey70/todo-app-go"
	"github.com/Shv-sergey70/todo-app-go/pkg/handler"
	"github.com/Shv-sergey70/todo-app-go/pkg/repository"
	"github.com/Shv-sergey70/todo-app-go/pkg/service"
	"log"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	server := new(todo.Server)

	if err := server.Run("8083", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server %s", err.Error())
	}
}
