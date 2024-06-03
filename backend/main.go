package main

import (
	"crud-api/api" // Use your module name here
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/users", api.CreateUser).Methods("POST")
	r.HandleFunc("/users/{name}", api.DeleteUser).Methods("DELETE")
	r.HandleFunc("/users/{name}/description", api.CreateDescription).Methods("POST")
	r.HandleFunc("/users/{name}/description", api.UpdateDescription).Methods("PUT")

	// Enable CORS
	headers := handlers.AllowedHeaders([]string{"Content-Type"})
	origins := handlers.AllowedOrigins([]string{"http://localhost:8081"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(headers, origins, methods)(r)))
}
