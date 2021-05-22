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
	uuidObj, err := uuid.NewUUID()
	if err != nil {
		log.Fatalln(err)
	}
	return uuidObj.String(), err
}

func StartMainServer() error {
	// gorilla/muxを使ったルーティング
	r := mux.NewRouter().StrictSlash(true)
	port := fmt.Sprintf(":%s", config.Config.Port)
	r.HandleFunc("/api/product", getAllProductsHandler).Methods("GET")
	r.HandleFunc("/api/product/{uuid}", getProductHandler).Methods("GET")
	r.HandleFunc("/api/accessory_category", getAllAccessoryCategoriesHandler).Methods("GET")
	r.HandleFunc("/api/accessory_category", createAccessoryCategoriesHandler).Methods("POST")
	r.HandleFunc("/api/material_category", getAllMaterialCategoriesHandler).Methods("GET")
	r.HandleFunc("/api/material_category", createMaterialCategoriesHandler).Methods("POST")
	r.HandleFunc("/api/sales_site", getAllSalesSitesHandler).Methods("GET")
	r.HandleFunc("/api/skill_market", getAllSkillMarketsHandler).Methods("GET")
	r.HandleFunc("/api/sns", getAllSnsListHandler).Methods("GET")
	r.HandleFunc("/api/creator", getCreatorHandler).Methods("GET")
	r.HandleFunc("/api/users", getAllUsersHandler).Methods("GET")
	// corsの設定
	c := cors.Default().Handler(r)
	return http.ListenAndServe(port, c)
}
