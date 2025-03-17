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

	r.HandleFunc("/tasks", controllers.GetTasks).Methods("GET")
	r.HandleFunc("/tasks", controllers.CreateTasks).Methods("GET")
	r.HandleFunc("/tasks", controllers.DeleteTasks).Methods("GET")
	r.HandleFunc("/tasks", controllers.UpdateTasks).Methods("GET")

	// Start the server
	fmt.Println("Server started on port 8085...")
	log.Fatal(http.ListenAndServe(":8085", r))
}
