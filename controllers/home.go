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

func loginHandler(w http.ResponseWriter, r *http.Request) {
	vop := views.LoginViewModelOp{}
	v := vop.GetVM()
	templates["login.html"].Execute(w, &v)
}
