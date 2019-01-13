package views

import "github.com/elfgzp/go_blog/models"

type ProfileEditViewModel struct {
	LoginViewModel
	ProfileUser models.User
}

type ProfileEditViewModelOp struct {
}

func (ProfileEditViewModelOp) GetVM(username string) ProfileEditViewModel {
	v := ProfileEditViewModel{}
	u, _ := models.GetUserByUsername(username)
	v.SetTitle("Profile Edit")
	v.SetCurrentUser(username)
	v.ProfileUser = *u
	return v
}

func UpdateAboutMe(username, text string) error {
	return models.UpdateAboutMe(username, text)
}
