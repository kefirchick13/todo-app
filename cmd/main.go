package main

import (
	todo "github.com/kefirchick13/todo-app"
	"log"
)

func main() {
	srv := new(todo.Server)
	if err := srv.Run("8000"); err != nil {
		log.Fatalf("Error during pushing: %s", err.Error())
	}

}
