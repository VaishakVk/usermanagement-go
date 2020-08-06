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
	userRouter.HandleFunc("", users.GetAllUsers).Methods(http.MethodGet)
	userRouter.HandleFunc("/login", users.LoginUser).Methods(http.MethodGet)
	userRouter.HandleFunc("/me", users.GetMe).Methods(http.MethodGet)
	userRouter.HandleFunc("/{id}", users.GetUserByID).Methods(http.MethodGet)
	userRouter.HandleFunc("/{id}", users.UpdateUser).Methods(http.MethodPut)
}
