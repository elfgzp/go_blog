package views

import (
	"github.com/elfgzp/go_blog/models"
	"log"
)

// LoginViewModel struct
type LoginViewModel struct {
	BaseViewModel
	Errs []string
}

// AddError func
func (v *LoginViewModel) AddError(errs ...string) {
	v.Errs = append(v.Errs, errs...)
}

// LoginViewModelOp struct
type LoginViewModelOp struct {
}

// GetVM func
func (LoginViewModelOp) GetVM() LoginViewModel {
	v := LoginViewModel{}
	v.SetTitle("Login")
	return v
}

func CheckLogin(username, password string) bool {
	user, err := models.GetUserByUsername(username)
	if err != nil {
		log.Println("Can not find username: ", username)
		log.Println("Error: ", err)
		return false
	}
	return user.CheckPassword(password)
}
