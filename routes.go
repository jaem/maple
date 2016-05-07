package maple

import (
	"github.com/jaem/maple/accounts"
	"github.com/gorilla/mux"
	"github.com/jaem/bounce"
	"github.com/jaem/bounce/storage/jwt"
	"github.com/jaem/nimble"
	"github.com/jaem/bounce/providers/local"
	"net/http"
	"fmt"
)

func initRoutes(web *nimble.Nimble) *mux.Router {
	bou := bounce.New(jwt.NewIdManager())
	bou.Register("local", local.NewProvider())
	web.UseHandlerFunc(bou.IdentifyRequest)

	// setup the routes
	router := mux.NewRouter()

	router.HandleFunc("/", helloHandlerFunc)
	router.HandleFunc("/hello", helloHandlerFunc).Methods("GET")

	buildAuthRoutes(router, bou)
	buildAppRoutes(router)

	return router
}


func helloHandlerFunc(w http.ResponseWriter, r *http.Request) {
	val := r.Header.Get("Cookie")
	fmt.Println("... val = "+ val)
	fmt.Fprintf(w, "Hello world!")
}

func buildAuthRoutes(router *mux.Router, bou *bounce.Bounce) {

	authRoutes := mux.NewRouter()
	authRoutes.HandleFunc("/auth/login", bou.Authenticate("local")).Methods("POST")
	authRoutes.HandleFunc("/auth/login", accounts.Login).Methods("GET")
	authRoutes.HandleFunc("/auth/register", accounts.Register)
	authRoutes.HandleFunc("/auth/verify", accounts.Verify)
	authRoutes.HandleFunc("/auth/recovery", accounts.Recovery)
	router.PathPrefix("/auth").Handler(nimble.New().
		Use(authRoutes))

	userRoutes := mux.NewRouter()
	authRoutes.HandleFunc("/user/logout", accounts.Logout)
	userRoutes.HandleFunc("/user/{userid:[0-9]+}/profile", accounts.ViewProfile)
	userRoutes.HandleFunc("/user/{userid:[0-9]+}/edit", accounts.ViewProfile)
	userRoutes.HandleFunc("/user/{userid:[0-9]+}/update", accounts.ViewProfile)
	router.PathPrefix("/user").Handler(nimble.New().
		UseHandlerFunc(bounce.IsLoggedIn).
		Use(userRoutes))
}

func buildAppRoutes(router *mux.Router) {

}
