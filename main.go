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
	server.Handle("/", http.FileServer(http.Dir("./public")))
	// server.Handle("/", http.FileServer(http.Dir(".")))

	// api endpoints
	server.HandleFunc("GET /tasks", routes.Tasks)

	fmt.Println("server up and running...")
	log.Fatal(http.ListenAndServe(PORT, server))
}
