package controllers

import (
	"api/app/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// アクセサリーカテゴリー一覧を取得
func getAllAccessoryCategoriesHandler(w http.ResponseWriter, r *http.Request) {
	accessoryCategories, err := models.GetAllAccessoryCategories()
	if err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("error: %s", err), http.StatusForbidden)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(accessoryCategories); err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("error: %s", err), http.StatusForbidden)
		return
	}
}

// アクセサリーカテゴリーの新規作成
func createAccessoryCategoryHandler(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("error: %s", err), http.StatusForbidden)
		return
	}

	var accessoryCategory models.AccessoryCategory
	if err := json.Unmarshal(reqBody, &accessoryCategory); err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("error: %s", err), http.StatusForbidden)
		return
	}
	// modelの呼び出し
	if err = models.InsertAccessoryCategory(&accessoryCategory); err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("error: %s", err), http.StatusForbidden)
		return
	}
	responseBody, err := json.Marshal(accessoryCategory)
	if err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("error: %s", err), http.StatusForbidden)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseBody)
}

// 材料カテゴリー一覧を取得
func getAllMaterialCategoriesHandler(w http.ResponseWriter, r *http.Request) {
	materialCategories, err := models.GetAllMaterialCategories()
	if err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("error: %s", err), http.StatusForbidden)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(materialCategories); err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("error: %s", err), http.StatusForbidden)
		return
	}
}

// 材料カテゴリーの新規作成
func createMaterialCategoryHandler(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("error: %s", err), http.StatusForbidden)
		return
	}

	var materialCategory models.MaterialCategory
	if err := json.Unmarshal(reqBody, &materialCategory); err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("error: %s", err), http.StatusForbidden)
		return
	}
	// modelの呼び出し
	models.InsertMaterialCategory(&materialCategory)
	responseBody, err := json.Marshal(materialCategory)
	if err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("error: %s", err), http.StatusForbidden)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseBody)
}
