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
	err := json.NewEncoder(w).Encode(accessoryCategories)
	if err != nil {
		log.Fatalln(err)
	}
}

// アクセサリーカテゴリーの新規作成
func createAccessoryCategoriesHandler(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	var accessory_category models.AccessoryCategory
	if err := json.Unmarshal(reqBody, &accessory_category); err != nil {
		log.Fatal(err)
	}
	// modelの呼び出し
	models.InsertAccessoryCategory(&accessory_category)
	responseBody, err := json.Marshal(accessory_category)
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
	err := json.NewEncoder(w).Encode(materialCategories)
	if err != nil {
		log.Fatalln(err)
	}
}

// 材料カテゴリーの新規作成
func createMaterialCategoriesHandler(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	var material_category models.MaterialCategory
	if err := json.Unmarshal(reqBody, &material_category); err != nil {
		log.Fatal(err)
	}
	// modelの呼び出し
	models.InsertMaterialCategory(&material_category)
	responseBody, err := json.Marshal(material_category)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseBody)
}