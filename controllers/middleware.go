package controllers

import (
	"fmt"
	"github.com/elfgzp/go_blog/models"
	"log"
	"net/http"
)

func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username, err := getSessionUser(r)
		log.Println(fmt.Printf("AuthMiddleware get username: %s from request: %s %s ", username, r.RemoteAddr, r.URL))

		if username != "" {
			log.Println(fmt.Sprintf("%s last seen update", username))
			_ = models.UpdateLastSeen(username)
		}

		if err != nil || len(username) == 0 {
			log.Println("AuthMiddleware get session error and redirect to login")
			http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
		} else {
			next.ServeHTTP(w, r)
		}
	}
}
