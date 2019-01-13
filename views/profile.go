package views

import "github.com/elfgzp/go_blog/models"

type ProfileViewModel struct {
	BaseViewModel
	Posts          []models.Post
	Editable       bool
	ProfileUser    models.User
	IsFollow       bool
	FollowersCount int
	FollowingCount int
}

type ProfileViewModelOp struct {
}

func (ProfileViewModelOp) GetVM(sUser, pUser string) (ProfileViewModel, error) {
	v := ProfileViewModel{}
	v.SetTitle("Profile")
	u, err := models.GetUserByUsername(pUser)
	if err != nil {
		return v, err
	}
	posts, _ := models.GetPostsByUserID(u.ID)
	v.ProfileUser = *u
	v.Editable = sUser == pUser

	if !v.Editable {
		v.IsFollow = u.IsFollowedByUser(sUser)
	}

	v.FollowersCount = u.FollowersCount()
	v.FollowingCount = u.FollowingCount()

	v.Posts = *posts
	v.SetCurrentUser(sUser)
	return v, nil
}

func Follow(sUser, pUser string) error {
	u, err := models.GetUserByUsername(sUser)
	if err != nil {
		return err
	}

	return u.Follow(pUser)
}

func UnFollow(sUser, pUser string) error {
	u, err := models.GetUserByUsername(sUser)
	if err != nil {
		return err
	}

	return u.UnFollow(pUser)
}
