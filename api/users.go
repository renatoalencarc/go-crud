package api

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)

type User struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

var users = make(map[string]User)
var mu sync.Mutex

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mu.Lock()
	users[user.Name] = user
	mu.Unlock()

	w.WriteHeader(http.StatusCreated)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	mu.Lock()
	delete(users, name)
	mu.Unlock()

	w.WriteHeader(http.StatusNoContent)
}
