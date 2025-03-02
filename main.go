package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/GOAT-stack-todo-app/routes"
)

func main() {

	fmt.Println("GOAT todo-app")

	// init the server.
	server := http.NewServeMux()
	PORT := ":5000"

	// static landing page
	// server.Handle("/", http.FileServer(http.Dir("./public")))

	// api endpoints
	server.HandleFunc("GET /", routes.HomePage)
	server.HandleFunc("POST /add-task", routes.AddTask)
	server.HandleFunc("GET /tasks", routes.GetAllTasks)
	server.HandleFunc("GET /finish/{id}", routes.FinishByID)
	server.HandleFunc("GET /delete/{id}", routes.DeleteByID)

	fmt.Println("server up and running...")
	log.Fatal(http.ListenAndServe(PORT, server))
}
