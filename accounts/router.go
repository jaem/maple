package accounts

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/jaem/nimble"
)

func BuildAuthRoutes(router *mux.Router) {

	authRoutes := mux.NewRouter()
	authRoutes.HandleFunc("/auth/login", login)
	authRoutes.HandleFunc("/auth/logout", logout)
	authRoutes.HandleFunc("/auth/register", register)
	authRoutes.HandleFunc("/auth/verify", verify)
	authRoutes.HandleFunc("/auth/recovery", recovery)

	router.PathPrefix("/auth").Handler(nimble.New().
		UseHandlerFunc(authMiddleware).Use(authRoutes))

	acctRoutes := mux.NewRouter()
	acctRoutes.HandleFunc("/account/profile", viewProfile)

	router.PathPrefix("/account").Handler(nimble.New().
		UseHandlerFunc(authMiddleware).Use(acctRoutes))
}

func authMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	// do some stuff before
	w.Write([]byte("......... in middleware AUTH"))
	next(w, r)
	// do some stuff after
}

func login(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("......... login"))
}

func logout(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("......... logout"))
}

func register(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("......... register"))
}

/**
  Verify user via email
 */
func verify(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("......... verify"))
}

func recovery(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("......... resetPassword"))
}

func viewProfile(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("......... viewProfile"))
}
