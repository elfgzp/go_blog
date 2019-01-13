package controllers

import (
	"fmt"
	"github.com/elfgzp/go_blog/views"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type home struct {
}

func (h home) registerRoutes() {
	r := mux.NewRouter()
	r.HandleFunc("/logout", authMiddleware(logoutHandler))
	r.HandleFunc("/login", loginHandler)
	r.HandleFunc("/register", registerHandler)
	r.HandleFunc("/user/{username}", authMiddleware(profileHandler))
	r.HandleFunc("/profile_edit", authMiddleware(profileEditHandler))
	r.HandleFunc("/follow/{username}", authMiddleware(followHandler))
	r.HandleFunc("/unFollow/{username}", authMiddleware(UnFollowHandler))
	r.HandleFunc("/", authMiddleware(indexHandler))

	http.Handle("/", r)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpName := "index.html"
	vop := views.IndexViewModelOp{}
	username, err := getSessionUser(r)
	page := getPage(r)

	if r.Method == http.MethodGet {
		flash := getFlash(w, r)
		if err != nil {
			panic(fmt.Errorf("IndexHandler getSessionUser error: %s", err))
		}

		v := vop.GetVM(username, flash, page, pageLimit)
		templates[tpName].Execute(w, &v)
	}

	if r.Method == http.MethodPost {
		_ = r.ParseForm()
		body := r.Form.Get("body")
		errMessage := checkLen("Post", body, 1, 180)
		if errMessage != "" {
			setFlash(w, r, errMessage)
		} else {
			err := views.CreatePost(username, body)
			if err != nil {
				log.Println("Add Post error: ", err)
				w.Write([]byte("Error insert Post in database"))
				return
			}
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	tpName := "login.html"
	vop := views.LoginViewModelOp{}
	v := vop.GetVM()
	if r.Method == http.MethodGet {
		templates[tpName].Execute(w, &v)
	}

	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			panic(fmt.Errorf("ParseForm error: %s", err))
		}

		username := r.Form.Get("username")
		password := r.Form.Get("password")

		errs := checkLogin(username, password)
		v.AddError(errs...)

		if len(v.Errs) > 0 {
			templates[tpName].Execute(w, &v)
		} else {
			err := setSessionUser(w, r, username)
			if err != nil {
				panic(fmt.Errorf("Set session failed with error: %s", err))
			}
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	}
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	tpName := "register.html"
	vop := views.RegisterViewModelOp{}
	v := vop.GetVM()
	if r.Method == http.MethodGet {
		templates[tpName].Execute(w, &v)
	}

	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			panic(fmt.Errorf("ParseForm error: %s", err))
		}

		username := r.Form.Get("username")
		email := r.Form.Get("email")
		pwd1 := r.Form.Get("pwd1")
		pwd2 := r.Form.Get("pwd2")

		errs := checkRegister(username, email, pwd1, pwd2)
		v.AddError(errs...)

		if len(v.Errs) > 0 {
			templates[tpName].Execute(w, &v)
		} else {
			if err := addUser(username, pwd1, email); err != nil {
				log.Println("add User error:", err)
				w.Write([]byte("Error insert database"))
				return
			}
			err := setSessionUser(w, r, username)
			if err != nil {
				panic(fmt.Errorf("Set session failed with error: %s", err))
			}
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	}
}

func profileHandler(w http.ResponseWriter, r *http.Request) {
	tpName := "profile.html"
	vars := mux.Vars(r)
	pUser := vars["username"]
	sUser, _ := getSessionUser(r)
	vop := views.ProfileViewModelOp{}
	v, err := vop.GetVM(sUser, pUser)
	if err != nil {
		msg := fmt.Sprintf("user ( %s ) does not exist", pUser)
		w.Write([]byte(msg))
		return
	}
	templates[tpName].Execute(w, &v)
}

func profileEditHandler(w http.ResponseWriter, r *http.Request) {
	tpName := "profile_edit.html"
	username, _ := getSessionUser(r)
	vop := views.ProfileEditViewModelOp{}
	v := vop.GetVM(username)
	if r.Method == http.MethodGet {
		err := templates[tpName].Execute(w, &v)
		if err != nil {
			log.Println(err)

		}
	}

	if r.Method == http.MethodPost {
		_ = r.ParseForm()
		aboutMe := r.Form.Get("aboutMe")
		if err := views.UpdateAboutMe(username, aboutMe); err != nil {
			log.Println(fmt.Sprintf("Update about me error: %s", err))
			w.Write([]byte("Error update about me"))
			return
		}

		http.Redirect(w, r, fmt.Sprintf("/user/%s", username), http.StatusSeeOther)
	}
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	err := clearSession(w, r)
	if err != nil {
		panic(fmt.Errorf("Clear session failed with error: %s", err))
	}
	http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
}

func followHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pUser := vars["username"]
	sUser, _ := getSessionUser(r)

	err := views.Follow(sUser, pUser)

	if err != nil {
		log.Println("Follow err: ", err)
		w.Write([]byte("Error in Follow"))
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/user/%s", pUser), http.StatusSeeOther)
}

func UnFollowHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pUser := vars["username"]
	sUser, _ := getSessionUser(r)

	err := views.UnFollow(sUser, pUser)

	if err != nil {
		log.Println("UnFollow err: ", err)
		w.Write([]byte("UnError in Follow"))
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/user/%s", pUser), http.StatusSeeOther)
}
