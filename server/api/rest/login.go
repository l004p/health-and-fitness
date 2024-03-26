package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	//"server/core/user"
	"server/services/authentication"

	"log"
)

type loginValues struct {
	Username	string `json:"username"`
	Password	string `json:"password"`
}

func (lh *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var loginValue loginValues
	err := json.NewDecoder(r.Body).Decode(&loginValue)
	if err != nil {
		log.Println(err)
		return
	}
	token, err := authentication.Login(lh.ur, r.Context(), loginValue.Username, loginValue.Password)
	if err != nil {
		log.Println(err)
		return
	}
	//jsonBytes, err := json.Marshal(token)
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, token)
	//w.Write(jsonBytes)
}