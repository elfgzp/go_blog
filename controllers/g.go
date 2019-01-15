package controllers

import (
	"github.com/gorilla/sessions"
	"html/template"
)

var (
	homeController home
	templates      map[string]*template.Template
	sessionName    string
	flashName      string
	store          *sessions.CookieStore
	pageLimit      int
)

func init() {
	templates = PopulateTemplates()
	store = sessions.NewCookieStore([]byte("someting-very-secret"))
	sessionName = "go_blog"
	flashName = "go-flash"
	pageLimit = 5
}

// Startup func
func Startup() {
	homeController.registerRoutes()
}

