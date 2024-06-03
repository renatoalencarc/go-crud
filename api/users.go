package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// User struct to hold the user data
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mu.Lock()
	users[user.Name] = user.Description
	mu.Unlock()

	w.WriteHeader(http.StatusCreated)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]

	mu.Lock()
	if _, exists := users[name]; exists {
		delete(users, name)
		mu.Unlock()
		w.WriteHeader(http.StatusNoContent)
	} else {
		mu.Unlock()
		http.Error(w, "User not found", http.StatusNotFound)
	}
}

func ListUsers(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	var userList []User
	for name, description := range users {
		userList = append(userList, User{Name: name, Description: description})
	}
	json.NewEncoder(w).Encode(userList)
}
