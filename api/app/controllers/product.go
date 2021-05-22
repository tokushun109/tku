package controllers

import (
	"api/app/models"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// 商品一覧を取得
func getAllProductsHandler(w http.ResponseWriter, r *http.Request) {
	products := models.GetAllProducts()
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(products); err != nil {
		log.Fatalln(err)
	}
}

// 商品詳細を取得
func getProductHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uuid := vars["uuid"]
	product := models.GetProduct(uuid)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(product); err != nil {
		log.Fatalln(err)
	}
}
