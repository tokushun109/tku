package controllers

import (
	"api/app/models"
	"encoding/json"
	"log"
	"net/http"
)

// アクセサリーカテゴリー一覧を取得
func getAllAccessoryCategoriesHandler(w http.ResponseWriter, r *http.Request) {
	accessoryCategories := models.GetAllAccessoryCategories()
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(accessoryCategories)
	if err != nil {
		log.Fatalln(err)
	}
}

// 材料カテゴリー一覧を取得
func getAllMaterialCategoriesHandler(w http.ResponseWriter, r *http.Request) {
	materialCategories := models.GetAllMaterialCategories()
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(materialCategories)
	if err != nil {
		log.Fatalln(err)
	}
}
