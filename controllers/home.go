package controllers

import (
	"github.com/elfgzp/go_blog/views"
	"net/http"
)

type home struct {
}

func (h home) registerRoutes() {
	http.HandleFunc("/", indexHandler)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	vop := views.IndexViewModelOp{}
	v := vop.GetVM()
	templates["index.html"].Execute(w, &v)
}
