package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateDescription(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	var description struct {
		Description string `json:"description"`
	}
	if err := json.NewDecoder(r.Body).Decode(&description); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mu.Lock()
	if user, exists := users[name]; exists {
		user.Description = description.Description
		users[name] = user
		w.WriteHeader(http.StatusCreated)
	} else {
		http.Error(w, "User not found", http.StatusNotFound)
	}
	mu.Unlock()
}

func UpdateDescription(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	var description struct {
		Description string `json:"description"`
	}
	if err := json.NewDecoder(r.Body).Decode(&description); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mu.Lock()
	if user, exists := users[name]; exists {
		user.Description = description.Description
		users[name] = user
		w.WriteHeader(http.StatusOK)
	} else {
		http.Error(w, "User not found", http.StatusNotFound)
	}
	mu.Unlock()
}
