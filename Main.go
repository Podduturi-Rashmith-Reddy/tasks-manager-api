package main

import (
	"fmt"
	"log"
	"net/http"

	"tasks-manager-api/controllers"
	"tasks-manager-api/database"

	"github.com/gorilla/mux"
)

func main() {
	// Connect to the database
	db := database.ConnectDB()
	defer db.Close()

	// Initialize the router
	r := mux.NewRouter()

	// Register API endpoints

	r.HandleFunc("/tasks", controllers.CreateTask).Methods("POST") // Create a new task

	// Start the server
	fmt.Println("Server started on port 8082...")
	log.Fatal(http.ListenAndServe(":8082", r))
}
