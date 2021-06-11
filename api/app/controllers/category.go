package controllers

import (
	"api/app/models"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// アクセサリーカテゴリー一覧を取得
func getAllAccessoryCategoriesHandler(w http.ResponseWriter, r *http.Request) {
	accessoryCategories := models.GetAllAccessoryCategories()
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(accessoryCategories); err != nil {
		log.Fatalln(err)
	}
}

// アクセサリーカテゴリーの新規作成
func createAccessoryCategoryHandler(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	var accessoryCategory models.AccessoryCategory
	if err := json.Unmarshal(reqBody, &accessoryCategory); err != nil {
		log.Fatal(err)
	}
	// modelの呼び出し
	models.InsertAccessoryCategory(&accessoryCategory)
	responseBody, err := json.Marshal(accessoryCategory)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseBody)
}

// 材料カテゴリー一覧を取得
func getAllMaterialCategoriesHandler(w http.ResponseWriter, r *http.Request) {
	materialCategories := models.GetAllMaterialCategories()
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(materialCategories); err != nil {
		log.Fatalln(err)
	}
}

// 材料カテゴリーの新規作成
func createMaterialCategoryHandler(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	var materialCategory models.MaterialCategory
	if err := json.Unmarshal(reqBody, &materialCategory); err != nil {
		log.Fatal(err)
	}
	// modelの呼び出し
	models.InsertMaterialCategory(&materialCategory)
	responseBody, err := json.Marshal(materialCategory)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseBody)
}
