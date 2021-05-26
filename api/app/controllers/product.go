package controllers

import (
	"api/app/models"
	"encoding/json"
	"io/ioutil"
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

// 商品の新規作成
func createProductHandler(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	var product models.Product
	if err := json.Unmarshal(reqBody, &product); err != nil {
		log.Fatal(err)
	}
	// uuidの設定
	uuid, err := GenerateUuid()
	if err != nil {
		log.Fatal(err)
	}
	product.Uuid = uuid
	// modelの呼び出し
	models.InsertProduct(&product)
	responseBody, err := json.Marshal(product)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(responseBody)
}
