package controllers

import (
	"api/config"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func StartMainServer() error {
	routing := mux.NewRouter().StrictSlash(true)
	port := fmt.Sprintf(":%s", config.Config.Port)
	routing.HandleFunc("/api/products", ProductsHandler).Methods("GET")
	return http.ListenAndServe(port, routing)
}
