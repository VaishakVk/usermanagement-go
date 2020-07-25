package routes

import (
	"goapi/lib/users"
	"net/http"

	"github.com/gorilla/mux"
)

// RegisterUserRoutes API
func RegisterUserRoutes(router *mux.Router) {
	userRouter := router.PathPrefix("/users").Subrouter()
	userRouter.HandleFunc("", users.CreateUser).Methods(http.MethodPost)
}
