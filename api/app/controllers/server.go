package controllers

import (
	"api/config"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type successResponse struct {
	Success bool `json:"success"`
}

// 処理の成功結果をレスポンスで返す
func getSuccessResponse() (responseBody []byte) {
	var successResponse successResponse
	successResponse.Success = true
	responseBody, _ = json.Marshal(successResponse)
	return responseBody
}

func StartMainServer() error {
	// gorilla/muxを使ったルーティング
	r := mux.NewRouter().StrictSlash(true)
	port := fmt.Sprintf(":%s", config.Config.Port)
	// ヘルスチェック
	r.HandleFunc("/api/health_check", getHealthCheckHandler).Methods("GET")
	// 商品
	r.HandleFunc("/api/product", getAllProductsHandler).Methods("GET")
	r.HandleFunc("/api/product/{product_uuid}", getProductHandler).Methods("GET")
	r.HandleFunc("/api/product", createProductHandler).Methods("POST")
	r.HandleFunc("/api/product/{product_uuid}", updateProductHandler).Methods("PUT")
	r.HandleFunc("/api/product/{product_uuid}", deleteProductHandler).Methods("DELETE")
	r.HandleFunc("/api/csv/product", getProductsCsvHandler).Methods("GET")
	r.HandleFunc("/api/csv/product", uploadProductsCsvHandler).Methods("POST")
	// 商品画像
	r.HandleFunc("/api/carousel_image", getCarouselImageHandler).Methods("GET")
	r.HandleFunc("/api/product_image/{product_image_uuid}/blob", getProductImageBlobHandler).Methods("GET")
	r.HandleFunc("/api/product/{product_uuid}/product_image", createProductImageHandler).Methods("POST")
	r.HandleFunc("/api/product/{product_uuid}/product_image/{product_image_uuid}", deleteProductImageHandler).Methods("DELETE")
	// カテゴリー
	r.HandleFunc("/api/category", getAllCategoriesHandler).Methods("GET")
	r.HandleFunc("/api/category", createCategoryHandler).Methods("POST")
	r.HandleFunc("/api/category/{category_uuid}", updateCategoryHandler).Methods("PUT")
	r.HandleFunc("/api/category/{category_uuid}", deleteCategoryHandler).Methods("DELETE")
	// タグ
	r.HandleFunc("/api/tag", getAllTagsHandler).Methods("GET")
	r.HandleFunc("/api/tag", createTagHandler).Methods("POST")
	r.HandleFunc("/api/tag/{tag_uuid}", updateTagHandler).Methods("PUT")
	r.HandleFunc("/api/tag/{tag_uuid}", deleteTagHandler).Methods("DELETE")
	// 販売サイト
	r.HandleFunc("/api/sales_site", getAllSalesSitesHandler).Methods("GET")
	r.HandleFunc("/api/sales_site", createSalesSiteHandler).Methods("POST")
	r.HandleFunc("/api/sales_site/{sales_site_uuid}", updateSalesSiteHandler).Methods("PUT")
	r.HandleFunc("/api/sales_site/{sales_site_uuid}", deleteSalesSiteHandler).Methods("DELETE")
	// スキルマーケット
	r.HandleFunc("/api/skill_market", getAllSkillMarketsHandler).Methods("GET")
	r.HandleFunc("/api/skill_market", createSkillMarketHandler).Methods("POST")
	r.HandleFunc("/api/skill_market/{skill_market_uuid}", updateSkillMarketHandler).Methods("PUT")
	r.HandleFunc("/api/skill_market/{skill_market_uuid}", deleteSkillMarketHandler).Methods("DELETE")
	// SNS
	r.HandleFunc("/api/sns", getAllSnsListHandler).Methods("GET")
	r.HandleFunc("/api/sns", createSnsHandler).Methods("POST")
	r.HandleFunc("/api/sns/{sns_uuid}", updateSnsHandler).Methods("PUT")
	r.HandleFunc("/api/sns/{sns_uuid}", deleteSnsHandler).Methods("DELETE")
	// 製作者
	r.HandleFunc("/api/creator", getCreatorHandler).Methods("GET")
	r.HandleFunc("/api/creator", updateCreatorHandler).Methods("PUT")
	r.HandleFunc("/api/creator/logo", updateCreatorLogoHandler).Methods("PUT")
	r.HandleFunc("/api/creator/logo/{logo_file}/blob", getCreatorLogoBlobHandler).Methods("GET")
	// お問い合わせ
	r.HandleFunc("/api/contact", getAllContactListHandler).Methods("GET")
	r.HandleFunc("/api/contact", createContactHandler).Methods("POST")
	// // ユーザー
	// r.HandleFunc("/api/user", getAllUsersHandler).Methods("GET")
	// ログイン
	r.HandleFunc("/api/user/login", getLoginUserHandler).Methods("GET")
	r.HandleFunc("/api/user/login", loginHandler).Methods("POST")
	// ログアウト
	r.HandleFunc("/api/user/logout", logoutHandler).Methods("POST")

	// corsの設定
	customizeCors := cors.New(cors.Options{
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
		AllowedOrigins:   []string{config.Config.ClientUrl},
		AllowCredentials: true,
	})
	c := customizeCors.Handler(r)
	return http.ListenAndServe(port, c)
}
