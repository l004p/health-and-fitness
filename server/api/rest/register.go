package rest

import (
	"net/http"
	"fmt"
)

func (lh *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)

	fmt.Fprintln(w, "hello")
}