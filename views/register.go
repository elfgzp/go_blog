package views

import (
	"github.com/elfgzp/go_blog/models"
	"log"
)

type RegisterViewModel struct {
	LoginViewModel
}

type RegisterViewModelOp struct {
}

func (RegisterViewModelOp) GetVM() RegisterViewModel {
	v := RegisterViewModel{}
	v.SetTitle("Register")
	return v
}

func CheckUserExist(username string) bool {
	_, err := models.GetUserByUsername(username)
	if err != nil {
		log.Println("Cant not find username: ", username)
		return true
	}
	return false
}

func AddUser(username, password, email string) error {
	return models.AddUser(username, password, email)
}
