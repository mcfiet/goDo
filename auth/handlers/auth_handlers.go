package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/mcfiet/goDo/user/model"
	"github.com/mcfiet/goDo/utils"
)

type AuthHandler struct{}

func (handler *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	log.Println("Login")
	log.Println("Header:", r.Header.Get("Authorization"))
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	log.Println(user)
	user.ID = uuid.New()

	if user.Username == "fiete" && user.Password == "test" {
		fmt.Println("Logged In")
		token, err := utils.GenerateToken(user.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Header().Set("Authorization", "Bearer "+token)
	} else {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
	}
}
