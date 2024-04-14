package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"server/core"
	"server/services/authentication"
	//userservice "server/services/user"
)

type registerValues struct {
	Username  string `json:"username"`
	UserEmail string `json:"user_email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
}

func (lh *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var registerValue registerValues
	err := json.NewDecoder(r.Body).Decode(&registerValue)
	if err != nil {
		log.Println(err)
		return
	}
	if registerValue.Username == "" || registerValue.UserEmail == "" || registerValue.FirstName == "" || registerValue.LastName == "" || registerValue.Password == "" {
		log.Println("all fields must be filled")
		return
	}

	user := core.User{
		Username: registerValue.Username,
		UserEmail: registerValue.UserEmail,
		FirstName: registerValue.FirstName,
		LastName: registerValue.LastName,
		UserPassword: registerValue.Password,
	}
	//err = lh.ur.FindOrCreate(r.Context(), user)
	//err = userservice.FindOrCreate(lh.ur, r.Context(), user)
	err = authentication.Register(lh.ur, r.Context(), user)
	if err != nil {
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "hello")
}