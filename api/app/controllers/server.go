package controllers

import (
	"api/config"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func GenerateUuid() (string, error) {
	uuidObj, err := uuid.NewRandom()
	if err != nil {
		log.Fatalln(err)
	}
	return uuidObj.String(), err
}

func StartMainServer() error {
	// gorilla/muxを使ったルーティング
	r := mux.NewRouter().StrictSlash(true)
	port := fmt.Sprintf(":%s", config.Config.Port)
	// 商品
	r.HandleFunc("/api/product", getAllProductsHandler).Methods("GET")
	r.HandleFunc("/api/product/{uuid}", getProductHandler).Methods("GET")
	r.HandleFunc("/api/product", createProductHandler).Methods("POST")
	// アクセサリーカテゴリー
	r.HandleFunc("/api/accessory_category", getAllAccessoryCategoriesHandler).Methods("GET")
	r.HandleFunc("/api/accessory_category", createAccessoryCategoriesHandler).Methods("POST")
	// 材料カテゴリー
	r.HandleFunc("/api/material_category", getAllMaterialCategoriesHandler).Methods("GET")
	r.HandleFunc("/api/material_category", createMaterialCategoriesHandler).Methods("POST")
	// 販売サイト
	r.HandleFunc("/api/sales_site", getAllSalesSitesHandler).Methods("GET")
	// スキルマーケット
	r.HandleFunc("/api/skill_market", getAllSkillMarketsHandler).Methods("GET")
	// SNS
	r.HandleFunc("/api/sns", getAllSnsListHandler).Methods("GET")
	// 製作者
	r.HandleFunc("/api/creator", getCreatorHandler).Methods("GET")
	// ユーザー
	r.HandleFunc("/api/users", getAllUsersHandler).Methods("GET")
	// corsの設定
	c := cors.Default().Handler(r)
	return http.ListenAndServe(port, c)
}
