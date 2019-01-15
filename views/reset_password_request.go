package views

import (
	"log"

	"github.com/elfgzp/go_blog/models"
)

type ResetPasswordRequestViewModel struct {
	LoginViewModel
}

type RestPasswordRequestViewModelOp struct {
}

func (RestPasswordRequestViewModelOp) GetVM() ResetPasswordRequestViewModel {
	v := ResetPasswordRequestViewModel{}
	v.SetTitle("Forget Password")
	return v
}

func CheckEmailExist(email string) bool {
	_, err := models.GetUserByEmail(email)
	if err != nil {
		log.Println("Can not find email:", email)
		return false
	}
	return true
}
