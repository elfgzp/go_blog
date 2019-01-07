package views

// LoginViewModel struct
type LoginViewModel struct {
	BaseViewModel
	Errs []string
}

// AddError func
func (v *LoginViewModel) AddError(errs ...string)  {
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
