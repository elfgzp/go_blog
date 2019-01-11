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
	http.HandleFunc("/logout", logoutHandler)
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
		err := r.ParseForm()
		if err != nil {
			panic(fmt.Errorf("ParseForm error: %s", err))
		}

		username := r.Form.Get("username")
		password := r.Form.Get("password")

		if len(username) < 3 {
			v.AddError("username must longer than 3")
		}

		if len(password) < 6 {
			v.AddError("password must longer than 6")
		}

		if !views.CheckLogin(username, password) {
			v.AddError("username or password not correct, please input again")
		}

		if len(v.Errs) > 0 {
			templates[tpName].Execute(w, &v)
		} else {
			err := setSessionUser(w, r, username)
			if err != nil {
				panic(fmt.Errorf("Set session failed with error: %s", err))
			}
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	}
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	err := clearSession(w, r)
	if err != nil {
		panic(fmt.Errorf("Clear session failed with error: %s", err))
	}
	http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
}
