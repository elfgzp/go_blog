package views

import (
	"github.com/elfgzp/go_blog/config"
	"github.com/elfgzp/go_blog/models"
)

type EmailViewModel struct {
	Username string
	Token    string
	Server   string
}

type EmailViewModelOp struct {
}

func (EmailViewModelOp) GetVM(email string) EmailViewModel {
	v := EmailViewModel{}
	u, _ := models.GetUserByEmail(email)
	v.Username = u.Username
	v.Token, _ = u.GenerateToken()
	v.Server = config.GetServerURL()
	return v
}
