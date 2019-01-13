package views

import "github.com/elfgzp/go_blog/models"

type ProfileViewModel struct {
	BaseViewModel
	Posts       []models.Post
	ProfileUser models.User
}

type ProfileViewModelOp struct {
}

func (ProfileViewModelOp) GetVM(sUser, pUser string) (ProfileViewModel, error) {
	v := ProfileViewModel{}
	v.SetTitle("Profile")
	u1, err := models.GetUserByUsername(pUser)
	if err != nil {
		return v, err
	}
	posts, _ := models.GetPostsByUserID(u1.ID)
	v.ProfileUser = *u1
	v.Posts = *posts
	v.SetCurrentUser(sUser)
	return v, nil
}
