package controllers

import (
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

func checkLogin(username, password string) bool {
	if username == "elfgzp" && password == "abc123" {
		return true
	}
	return false
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

		if len(username) < 3 {
			v.AddError("username must longer than 3")
		}

		if len(password) < 6 {
			v.AddError("password must longer than 6")
		}

		if !checkLogin(username, password) {
			v.AddError("username or password not correct, please input again")
		}

		if len(v.Errs) > 0 {
			templates[tpName].Execute(w, &v)
		} else {
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	}
}
