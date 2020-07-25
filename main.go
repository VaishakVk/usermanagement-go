package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"goapi/db"
	"goapi/middlewares"
	"goapi/routes"
)

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello")
}
func main() {
	fmt.Println("Hello W")
	db.Connect()
	router := mux.NewRouter()
	routes.RegisterUserRoutes(router)
	urlUpdated := middlewares.RemoveTrailingSlashes(router)
	headersAdded := middlewares.SetHeaders(urlUpdated)
	http.ListenAndServe(":8000", headersAdded)
}
