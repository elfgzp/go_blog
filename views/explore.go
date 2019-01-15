package views

import "github.com/elfgzp/go_blog/models"

type ExploreViewModel struct {
	BaseViewModel
	Posts []models.Post
	BasePageViewModel
}

type ExploreViewModelOp struct {
}

func (ExploreViewModelOp) GetVM(username string, page, limit int) ExploreViewModel {
	posts, total, _ := models.GetPostsByPageAndLimit(page, limit)

	v := ExploreViewModel{}
	v.SetTitle("Explore")
	v.Posts = *posts
	v.SetBasePageViewModel(total, page, limit)
	v.SetCurrentUser(username)

	return v
}
