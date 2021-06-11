package controllers

import (
	"api/config"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func StartMainServer() error {
	// gorilla/muxを使ったルーティング
	r := mux.NewRouter().StrictSlash(true)
	port := fmt.Sprintf(":%s", config.Config.Port)
	// 商品
	r.HandleFunc("/api/product", getAllProductsHandler).Methods("GET")
	r.HandleFunc("/api/product/{product_uuid}", getProductHandler).Methods("GET")
	r.HandleFunc("/api/product", createProductHandler).Methods("POST")
	// 商品画像
	r.HandleFunc("/api/product_image/{product_image_uuid}/blob", getProductImageBlobHandler).Methods("GET")
	r.HandleFunc("/api/product/{product_uuid}/product_image", createProductImageHandler).Methods("POST")
	// アクセサリーカテゴリー
	r.HandleFunc("/api/accessory_category", getAllAccessoryCategoriesHandler).Methods("GET")
	r.HandleFunc("/api/accessory_category", createAccessoryCategoryHandler).Methods("POST")
	// 材料カテゴリー
	r.HandleFunc("/api/material_category", getAllMaterialCategoriesHandler).Methods("GET")
	r.HandleFunc("/api/material_category", createMaterialCategoryHandler).Methods("POST")
	// 販売サイト
	r.HandleFunc("/api/sales_site", getAllSalesSitesHandler).Methods("GET")
	r.HandleFunc("/api/sales_site", createSalesSiteHandler).Methods("POST")
	// スキルマーケット
	r.HandleFunc("/api/skill_market", getAllSkillMarketsHandler).Methods("GET")
	r.HandleFunc("/api/skill_market", createSkillMarketHandler).Methods("POST")
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
