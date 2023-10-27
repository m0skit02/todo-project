package main

import (
	todo_project "github.com/m0skit02/todo-project"
	"github.com/m0skit02/todo-project/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)

	srv := new(todo_project.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("errir occured while running http server: %s", err.Error())
	}
}
