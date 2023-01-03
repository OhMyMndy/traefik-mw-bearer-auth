package main

import (
	"net/http"
	"os"
	"strings"
)

func loadToken() {
	tokenString := os.Getenv("TOKENS")
	for _, v := range strings.Split(tokenString, ",") {
		if len(v) > 0 {
			tokens[v] = true
		}
	}
}

var tokens = make(map[string]bool)

func validToken(token string) bool {
	_, ok := tokens[token]
	return ok
}

func handler(w http.ResponseWriter, r *http.Request) {
	authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")

	if len(authHeader) != 2 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if validToken(authHeader[1]) {
		w.WriteHeader(http.StatusOK)
		return
	}
	w.WriteHeader(http.StatusUnauthorized)

}
func main() {
	loadToken()
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8087", nil)
}
