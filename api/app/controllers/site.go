package controllers

import (
	"api/app/models"
	"encoding/json"
	"log"
	"net/http"
)

// 販売サイト一覧を取得
func getAllSalesSitesHandler(w http.ResponseWriter, r *http.Request) {
	salesSites := models.GetAllSalesSites()
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(salesSites)
	if err != nil {
		log.Fatalln(err)
	}
}

// スキルマーケット一覧を取得
func getAllSkillMarketsHandler(w http.ResponseWriter, r *http.Request) {
	skillMarkets := models.GetAllSkillMarkets()
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(skillMarkets)
	if err != nil {
		log.Fatalln(err)
	}
}

// SNS一覧を取得
func getAllSnsListHandler(w http.ResponseWriter, r *http.Request) {
	snsList := models.GetAllSnsList()
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(snsList)
	if err != nil {
		log.Fatalln(err)
	}
}
