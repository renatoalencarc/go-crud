package api

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)

// User struct to hold the user data
var (
	users = make(map[string]string)
	mu    sync.Mutex
)

type User struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func CreateDescription(w http.ResponseWriter, r *http.Request) {
	var desc struct {
		Description string `json:"description"`
	}
	name := mux.Vars(r)["name"]
	if err := json.NewDecoder(r.Body).Decode(&desc); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()
	if _, exists := users[name]; exists {
		users[name] = desc.Description
		w.WriteHeader(http.StatusOK)
	} else {
		http.Error(w, "User not found", http.StatusNotFound)
	}
}

func UpdateDescription(w http.ResponseWriter, r *http.Request) {
	var desc struct {
		Description string `json:"description"`
	}
	name := mux.Vars(r)["name"]
	if err := json.NewDecoder(r.Body).Decode(&desc); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()
	if _, exists := users[name]; exists {
		users[name] = desc.Description
		w.WriteHeader(http.StatusOK)
	} else {
		http.Error(w, "User not found", http.StatusNotFound)
	}
}
