package accounts

import (
	"net/http"
	"html/template"
)

var templates = template.Must(template.ParseFiles(
	"views/auth_login.tpl",
	//"/views/auth_status.tpl",
))

func AuthMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	// do some stuff before
	w.Write([]byte("......... in middleware AUTH"))
	next(w, r)
	// do some stuff after
}

func Login(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "auth_login.tpl", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("......... logout"))
}

func Register(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("......... register"))


	//user := accounts.User{ Name: "slene", Password: "hello" }
	//accounts.CreateUser(&user)

}

/**
  Verify user via email
 */
func Verify(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("......... verify"))
}

func Recovery(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("......... resetPassword"))
}

func ViewProfile(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("......... viewProfile"))
}
