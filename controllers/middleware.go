package controllers

import (
	"fmt"
	"log"
	"net/http"
)

func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username, err := getSessionUser(r)
		log.Println(fmt.Printf("AuthMiddleware get username: %s from request: %s %s ", username, r.RemoteAddr, r.URL))
		if err != nil || len(username) == 0 {
			log.Println("AuthMiddleware get session error and redirect to login")
			http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
		} else {
			next.ServeHTTP(w, r)
		}
	}
}
