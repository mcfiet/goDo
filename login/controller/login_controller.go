package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/mcfiet/goDo/login/model"
)

var store = sessions.NewCookieStore([]byte("secret"))

func GetLogin(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement
	fmt.Println("Login GET")
	session, err := store.Get(r, "session-name")
	if err != nil {
		http.Error(w, "Konnte Session nicht erstellen", http.StatusInternalServerError)
	}

	fmt.Println(session.Values)
}

func Login(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement
	fmt.Println("Login POST")
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	fmt.Println(user)

	if user.Username == "fiete" && user.Password == "test" {
		fmt.Println("Logged In")
		session, err := store.Get(r, "session-name")
		if err != nil {
			http.Error(w, "Konnte Session nicht erstellen", http.StatusInternalServerError)
		}

		session.Values["user"] = user.Username
		if err := session.Save(r, w); err != nil {
			http.Error(w, "Konnte Session nicht speichern", http.StatusInternalServerError)
			return
		}
	} else {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
	}

}
