package views

// BaseView struct
type BaseViewModel struct {
	Title string
}

// SetTitle func
func (v *BaseViewModel) SetTitle(title string) {
	v.Title = title
}
