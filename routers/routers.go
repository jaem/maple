package routers

import (
	"github.com/jaem/maple/accounts"
	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter()
	accounts.BuildAuthRoutes(router)
	return buildAppRoutes(router)
}

func buildAppRoutes(router *mux.Router) *mux.Router {
	return router
}
