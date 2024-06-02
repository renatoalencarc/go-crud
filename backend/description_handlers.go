package backend

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetUserDescription(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range users {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item.Description)
			return
		}
	}
	http.Error(w, "User not found", http.StatusNotFound)
}

func UpdateUserDescription(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range users {
		if item.ID == params["id"] {
			var user User
			json.NewDecoder(r.Body).Decode(&user)
			users[index].Description = user.Description
			json.NewEncoder(w).Encode(users[index])
			return
		}
	}
	http.Error(w, "User not found", http.StatusNotFound)
}
