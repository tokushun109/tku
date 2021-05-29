package controllers

import (
	"api/app/models"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// 販売サイト一覧を取得
func getAllSalesSitesHandler(w http.ResponseWriter, r *http.Request) {
	salesSites := models.GetAllSalesSites()
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(salesSites); err != nil {
		log.Fatalln(err)
	}
}

// 販売サイトの新規作成
func createSalesSitesHandler(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	var salesSite models.SalesSite
	if err := json.Unmarshal(reqBody, &salesSite); err != nil {
		log.Fatal(err)
	}
	// modelの呼び出し
	models.InsertSalesSite(&salesSite)
	responseBody, err := json.Marshal(salesSite)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseBody)
}

// スキルマーケット一覧を取得
func getAllSkillMarketsHandler(w http.ResponseWriter, r *http.Request) {
	skillMarkets := models.GetAllSkillMarkets()
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(skillMarkets); err != nil {
		log.Fatalln(err)
	}
}

// SNS一覧を取得
func getAllSnsListHandler(w http.ResponseWriter, r *http.Request) {
	snsList := models.GetAllSnsList()
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(snsList); err != nil {
		log.Fatalln(err)
	}
}
