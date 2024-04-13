package rest

import (
	"net/http"
	"fmt"
	"encoding/json"
	"log"
	"server/core"

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
	user := core.User{
		Username: registerValue.Username,
		UserEmail: registerValue.UserEmail,
		FirstName: registerValue.FirstName,
		LastName: registerValue.LastName,
		UserPassword: registerValue.Password,
	}
	err = lh.ur.FindOrCreate(r.Context(), user)
	if err != nil {
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "hello")
}