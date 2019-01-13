package views

import "github.com/elfgzp/go_blog/models"

// IndexViewModel struct
type IndexViewModel struct {
	BaseViewModel
	Posts []models.Post
	Flash string

	BasePageViewModel
}

// IndexViewModelOp struct
type IndexViewModelOp struct {
}

// GetVM func
func (IndexViewModelOp) GetVM(username, flash string, page, limit int) IndexViewModel {
	u, _ := models.GetUserByUsername(username)
	posts, total, _ := u.FollowingPostsByPageAndLimit(page, limit)
	v := IndexViewModel{}
	v.SetTitle("Homepage")
	v.Posts = *posts
	v.Flash = flash
	v.SetBasePageViewModel(total, page, limit)
	v.SetCurrentUser(username)
	return v
}

func CreatePost(username, body string) error {
	u, _ := models.GetUserByUsername(username)
	return u.CreatePost(body)
}
