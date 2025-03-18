package main

import (
	"fmt"
	"log"
	"net/http"
	"task-manager/controllers"
	"task-manager/database"

	"github.com/gorilla/mux"
)

func main() {
	// Connect to database
	db := database.ConnectDB()
	defer db.Close()

	// Create a new router
	r := mux.NewRouter()

	// Define routes
	r.HandleFunc("/tasks", controllers.GetTasks).Methods("GET")
	r.HandleFunc("/tasks", controllers.CreateTask).Methods("POST")
	r.HandleFunc("/tasks/{id}", controllers.UpdateTask).Methods("PUT")
	r.HandleFunc("/tasks/{id}", controllers.DeleteTask).Methods("DELETE")

	// Start server
	fmt.Println("Server running on port 8081...")
	log.Fatal(http.ListenAndServe(":8081", r))
}
