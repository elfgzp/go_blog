package views

import (
	"github.com/elfgzp/go_blog/models"
)

type ResetPasswordViewModel struct {
	LoginViewModel
	Token string
}

type ResetPasswordViewModelOp struct {
}

func (ResetPasswordViewModelOp) GetVM(token string) ResetPasswordViewModel {
	v := ResetPasswordViewModel{}
	v.SetTitle("Reset Password")
	v.Token = token
	return v
}

func CheckToken(tokenString string) (string, error) {
	return models.CheckToken(tokenString)
}

func ResetUserPassword(username, password string) error {
	return models.UpdatePassword(username, password)
}
