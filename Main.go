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
	r.HandleFunc("/tasks", controllers.GetTasks).Methods("GET")           // Retrieve all tasks
	r.HandleFunc("/tasks", controllers.CreateTask).Methods("POST")        // Create a new task
	r.HandleFunc("/tasks/{id}", controllers.UpdateTask).Methods("PUT")    // Update an existing task (including status)
	r.HandleFunc("/tasks/{id}", controllers.DeleteTask).Methods("DELETE") // Mark a task as deleted (update status)

	// Start the server
	fmt.Println("Server started on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
