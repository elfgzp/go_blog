package controllers

import (
	"fmt"
	"github.com/elfgzp/go_blog/views"
	"net/http"
)

type home struct {
}

func (h home) registerRoutes() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/login", loginHandler)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	vop := views.IndexViewModelOp{}
	v := vop.GetVM()
	templates["index.html"].Execute(w, &v)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	tpName := "login.html"
	vop := views.LoginViewModelOp{}
	v := vop.GetVM()
	if r.Method == http.MethodGet {
		templates[tpName].Execute(w, &v)
	}

	if r.Method == http.MethodPost {
		r.ParseForm()
		username := r.Form.Get("username")
		password := r.Form.Get("password")
		fmt.Fprintf(w, "Username: %s Password: %s", username, password)
	}
}
