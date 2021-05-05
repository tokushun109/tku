package controllers

import (
	"api/config"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func StartMainServer() error {
	r := mux.NewRouter().StrictSlash(true)
	port := fmt.Sprintf(":%s", config.Config.Port)
	r.HandleFunc("/api/product", getAllProductsHandler).Methods("GET")
	r.HandleFunc("/api/product/{uuid}", getProductHandler).Methods("GET")
	return http.ListenAndServe(port, r)
}
