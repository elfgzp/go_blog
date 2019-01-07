package views

// LoginViewModel struct
type LoginViewModel struct {
	BaseViewModel
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
