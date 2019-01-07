package views

import "github.com/elfgzp/go_blog/models"

// IndexViewModel struct
type IndexViewModel struct {
	BaseViewModel
	models.User
	Posts []models.Post
}

// IndexViewModelOp struct
type IndexViewModelOp struct {
}

// GetVM func
func (IndexViewModelOp) GetVM() IndexViewModel {
	u1 := models.User{Username: "elfgzp"}
	u2 := models.User{Username: "Tom"}

	posts := []models.Post{
		{User: u1, Body: "Beautiful day in Portland!"},
		{User: u2, Body: "The Avengers movie was so cool!"},
	}

	v := IndexViewModel{BaseViewModel{Title: "Homepage"}, u1, posts}
	return v
}
