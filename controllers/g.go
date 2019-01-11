package controllers

import (
	"github.com/gorilla/sessions"
	"html/template"
)

var (
	homeController home
	templates      map[string]*template.Template
	sessionName    string
	store          *sessions.CookieStore
)

func init() {
	templates = PopulateTemplates()
	store = sessions.NewCookieStore([]byte("someting-very-secret"))
	sessionName = "go_blog"
}

// Startup func
func Startup() {
	homeController.registerRoutes()
}
