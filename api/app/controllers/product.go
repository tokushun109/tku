package controllers

import (
	"api/app/models"
	"encoding/json"
	"log"
	"net/http"
)

// 商品一覧を取得
func getProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	products, err := models.GetProducts()
	if err != nil {
		log.Println(err)
	}
	json.NewEncoder(w).Encode(products)
}
