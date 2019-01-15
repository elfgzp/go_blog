package controllers

import (
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/elfgzp/go_blog/config"
	"github.com/elfgzp/go_blog/views"
	"gopkg.in/gomail.v2"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
)

// PopulateTemplates func
// Create map template name to template.Template
func PopulateTemplates() map[string]*template.Template {
	const basePath = "templates"
	result := make(map[string]*template.Template)
	layout := template.Must(template.ParseFiles(basePath + "/_base.html"))
	dir, err := os.Open(basePath + "/content")

	if err != nil {
		panic("Failed to open tempalte blocks directory: " + err.Error())
	}

	fis, err := dir.Readdir(-1)

	if err != nil {
		panic("Failed to read contents of content directory: " + err.Error())
	}

	for _, fi := range fis {
		func() {
			f, err := os.Open(basePath + "/content/" + fi.Name())

			if err != nil {
				panic("Failed to open template '" + fi.Name() + "'")
			}

			defer f.Close()

			content, err := ioutil.ReadAll(f)

			if err != nil {
				panic("Failed to read content from file '" + fi.Name() + "'")
			}

			tmpl := template.Must(layout.Clone())
			_, err = tmpl.Parse(string(content))

			if err != nil {
				panic("Failed to parse contents of '" + fi.Name() + "' as template")
			}

			result[fi.Name()] = tmpl
		}()
	}

	return result
}

// session
func getSessionUser(r *http.Request) (string, error) {
	var username string
	session, err := store.Get(r, sessionName)
	if err != nil {
		return "", err
	}
	val := session.Values["user"]
	fmt.Println("val:", val)
	username, ok := val.(string)
	if !ok {
		return "", errors.New("Can not get session user")
	}
	fmt.Println("username:", username)
	return username, nil
}

// setSessionUser
func setSessionUser(w http.ResponseWriter, r *http.Request, username string) error {
	session, err := store.Get(r, sessionName)
	if err != nil {
		return err
	}
	session.Values["user"] = username
	err = session.Save(r, w)
	if err != nil {
		return err
	}
	return nil
}

// clearSession
func clearSession(w http.ResponseWriter, r *http.Request) error {
	session, err := store.Get(r, sessionName)
	if err != nil {
		return err
	}

	session.Options.MaxAge = -1

	err = session.Save(r, w)

	if err != nil {
		return err
	}

	return nil
}

func setFlash(w http.ResponseWriter, r *http.Request, message string) {
	session, _ := store.Get(r, sessionName)
	session.AddFlash(message, flashName)
	_ = session.Save(r, w)
}

func getFlash(w http.ResponseWriter, r *http.Request) string {
	session, _ := store.Get(r, sessionName)
	fm := session.Flashes(flashName)
	if fm == nil {
		return ""
	}

	_ = session.Save(r, w)
	return fmt.Sprintf("%v", fm[0])
}

// Check functions
func checkLen(fieldName, fieldValue string, minLen, maxLen int) string {
	lenField := len(fieldValue)

	if lenField < minLen {
		return fmt.Sprintf("%s field is too short, less than %d.", fieldName, minLen)
	}

	if lenField > maxLen {
		return fmt.Sprintf("%s field is too long, more than %d.", fieldName, maxLen)
	}

	return ""
}

func checkUsername(username string) string {
	return checkLen("Username", username, 3, 20)
}

func checkPassword(password string) string {
	return checkLen("Password", password, 6, 50)
}

func checkEmail(email string) string {
	if len(email) == 0 {
		return fmt.Sprintf("Email field is required.")
	}

	if m, _ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4})$`, email); !m {
		return fmt.Sprintf("Email field is not a valid email.")
	}
	return ""
}

func checkUserPassword(username, password string) string {
	if !views.CheckLogin(username, password) {
		return fmt.Sprintf("Username or password is not correct.")
	}
	return ""
}

func checkUserExist(username string) string {
	if !views.CheckUserExist(username) {
		return fmt.Sprintf("Username already exist, please choose another username.")
	}
	return ""
}

func checkPwdRepeatMatch(pwd1, pwd2 string) string {
	if pwd1 != pwd2 {
		return fmt.Sprintf("2 password does not match.")
	}
	return ""
}

// checkLogin()

func checkLogin(username, password string) []string {
	var errs []string
	if errCheck := checkUsername(username); len(errCheck) > 0 {
		errs = append(errs, errCheck)
	}

	if errCheck := checkPassword(password); len(errCheck) > 0 {
		errs = append(errs, errCheck)
	}

	if errCheck := checkUserPassword(username, password); len(errCheck) > 0 {
		errs = append(errs, errCheck)
	}

	return errs
}

func checkRegister(username, email, pwd1, pwd2 string) []string {
	var errs []string
	if errCheck := checkUsername(username); len(errCheck) > 0 {
		errs = append(errs, errCheck)
	}

	if errCheck := checkPassword(pwd1); len(errCheck) > 0 {
		errs = append(errs, errCheck)
	}

	if errCheck := checkPwdRepeatMatch(pwd1, pwd2); len(errCheck) > 0 {
		errs = append(errs, errCheck)
	}

	if errCheck := checkEmail(email); len(errCheck) > 0 {
		errs = append(errs, errCheck)
	}
	if errCheck := checkUserExist(username); len(errCheck) > 0 {
		errs = append(errs, errCheck)
	}
	return errs
}

// addUser func
func addUser(username, password, email string) error {
	return views.AddUser(username, password, email)
}

func getPage(r *http.Request) int {
	url := r.URL
	query := url.Query()

	q := query.Get("page")
	if q == "" {
		return 1
	}

	page, err := strconv.Atoi(q)
	if err != nil {
		return 1
	}

	return page
}

func sendEmail(target, subject, content string) {
	server, port, usr, pwd := config.GetSMTPConfig()
	d := gomail.NewDialer(server, port, usr, pwd)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	m := gomail.NewMessage()
	m.SetHeader("From", usr)
	m.SetHeader("To", target)
	m.SetAddressHeader("Cc",  usr, "admin")
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", content)

	if err := d.DialAndSend(m); err != nil {
		log.Println("Email Error:", err)
		return
	}
}

func checkResetPasswordRequest(email string) []string {
	var errs []string
	exits := views.CheckEmailExist(email)
	if !exits {
		errs = append(errs, "Can not find email: ", email )
	}
	return errs
}

func checkResetPassword(pwd1, pwd2 string) []string  {
	var errs []string
	err := checkPwdRepeatMatch(pwd1, pwd2)
	if err != "" {
		errs = append(errs, err)
	}
	return errs
}