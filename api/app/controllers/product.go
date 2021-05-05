package controllers

import (
	"api/app/models"
	"encoding/json"
	"net/http"
)

// 商品一覧を取得
func getProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	products := models.GetProducts()
	json.NewEncoder(w).Encode(products)
}
