package main

import (
	"fmt"
	"net/http"

	"github.com/rs/cors"

	"crud-api/backend"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter() // Create a new router

	userRouter := router.PathPrefix("/users").Subrouter()
	userRouter.HandleFunc("/", backend.CreateUser).Methods("POST")
	userRouter.HandleFunc("/", backend.GetAllUsers).Methods("GET")
	userRouter.HandleFunc("/{id}", backend.DeleteUser).Methods("DELETE")

	descriptionRouter := router.PathPrefix("/descriptions").Subrouter()
	descriptionRouter.HandleFunc("/{id}", backend.GetUserDescription).Methods("GET")
	descriptionRouter.HandleFunc("/{id}", backend.UpdateUserDescription).Methods("PUT")

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./")))

	fmt.Println("Server is running on port 8081")

	// This will ONLY allow requests from the frontend origin at localhost
	handler := cors.New(cors.Options{
		AllowedOrigins: []string{"http://127.0.0.1:8080"}, // Frontend origin for test purposes
	}).Handler(router)
	http.ListenAndServe(":8081", handler)
}
