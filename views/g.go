package views

import "log"

// BaseView struct
type BaseViewModel struct {
	Title       string
	CurrentUser string
}

// SetTitle func
func (v *BaseViewModel) SetTitle(title string) {
	v.Title = title
}

// SetCurrentUser func
func (v *BaseViewModel) SetCurrentUser(username string) {
	v.CurrentUser = username
}

type BasePageViewModel struct {
	PrevPage    int
	NextPage    int
	Total       int
	CurrentPage int
	Limit       int
}

func (v *BasePageViewModel) SetPrevAndNextPage() {
	if v.CurrentPage > 1 {
		v.PrevPage = v.CurrentPage - 1
	}

	if (v.Total-1)/v.Limit >= v.CurrentPage {
		v.NextPage = v.CurrentPage + 1
		log.Println(v.NextPage)
	}
}

func (v *BasePageViewModel) SetBasePageViewModel(total, page, limit int) {
	v.Total = total
	v.CurrentPage = page
	v.Limit = limit
	v.SetPrevAndNextPage()
}
