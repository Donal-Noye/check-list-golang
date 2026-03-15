package main

import (
	"check-list/db"
	"check-list/http"
	"check-list/todo"
	"fmt"
)

func main() {
	db.CheckConnection()

	todoList := todo.GetList()
	httpHandlers := http.NewHTTPHandlers(todoList)
	httpServer := http.NewHTTPServer(httpHandlers)

	if err := httpServer.StartServer(); err != nil {
		fmt.Println("Error starting HTTP server:", err)
	}
}
