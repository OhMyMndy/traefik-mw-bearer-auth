package main

import (
	"fmt"
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
	
	fmt.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
	
	status := http.StatusUnauthorized
	
	if len(authHeader) != 2 {
		status = http.StatusBadRequest
	} else if validToken(authHeader[1]) {
		status = http.StatusOK
	}
	
	fmt.Printf("%s %s %s %s\n", r.RemoteAddr, r.Method, r.URL, status)
	
	w.WriteHeader(status)
}

func main() {
	httpPort := 8087
	loadToken()
	http.HandleFunc("/", handler)
	
	fmt.Printf("listening on %v\n", httpPort)
	
	http.ListenAndServe(fmt.Sprintf(":%v", httpPort), nil)
}
