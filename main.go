package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"goapi/db"
	"goapi/middlewares"
	"goapi/routes"
	"goapi/validator"
)

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello")
}
func main() {
	fmt.Println("Hello W")
	db.Connect()
	router := mux.NewRouter()
	validator.Init()
	routes.RegisterUserRoutes(router)
	errorsHandled := middlewares.Recovery(router)
	urlUpdated := middlewares.RemoveTrailingSlashes(errorsHandled)
	headersAdded := middlewares.SetHeaders(urlUpdated)
	http.ListenAndServe(":8000", headersAdded)
	// defer func() {
	// 	if r := recover(); r != nil {
	// 		fmt.Println("Inside Recover..")
	// 	}
	// }()
}
