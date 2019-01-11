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
func (IndexViewModelOp) GetVM(username string) IndexViewModel {
	u1, _ := models.GetUserByUsername(username)
	posts, _ := models.GetPostsByUserID(u1.ID)
	v := IndexViewModel{BaseViewModel{Title: "Homepage"}, *u1, *posts}
	v.SetCurrentUser(username)
	return v
}
