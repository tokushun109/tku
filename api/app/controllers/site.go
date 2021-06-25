package controllers

import (
	"api/app/models"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// 販売サイト一覧を取得
func getAllSalesSitesHandler(w http.ResponseWriter, r *http.Request) {
	salesSites, err := models.GetAllSalesSites()
	if err != nil {
		ErrorHandler(w, err, http.StatusForbidden)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(salesSites); err != nil {
		ErrorHandler(w, err, http.StatusForbidden)
		return
	}
}

// 販売サイトの新規作成
func createSalesSiteHandler(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		ErrorHandler(w, err, http.StatusForbidden)
		return
	}

	var salesSite models.SalesSite
	if err := json.Unmarshal(reqBody, &salesSite); err != nil {
		ErrorHandler(w, err, http.StatusForbidden)
		return
	}
	// modelの呼び出し
	err = models.InsertSalesSite(&salesSite)
	if err != nil {
		ErrorHandler(w, err, http.StatusForbidden)
		return
	}
	responseBody, err := json.Marshal(salesSite)
	if err != nil {
		ErrorHandler(w, err, http.StatusForbidden)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseBody)
}

// スキルマーケット一覧を取得
func getAllSkillMarketsHandler(w http.ResponseWriter, r *http.Request) {
	skillMarkets, err := models.GetAllSkillMarkets()
	if err != nil {
		ErrorHandler(w, err, http.StatusForbidden)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(skillMarkets); err != nil {
		ErrorHandler(w, err, http.StatusForbidden)
		return
	}
}

// スキルマーケットの新規作成
func createSkillMarketHandler(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		ErrorHandler(w, err, http.StatusForbidden)
		return
	}

	var skillmarket models.SkillMarket
	if err := json.Unmarshal(reqBody, &skillmarket); err != nil {
		ErrorHandler(w, err, http.StatusForbidden)
		return
	}
	// modelの呼び出し
	err = models.InsertSkillMarket(&skillmarket)
	if err != nil {
		ErrorHandler(w, err, http.StatusForbidden)
		return
	}
	responseBody, err := json.Marshal(skillmarket)
	if err != nil {
		ErrorHandler(w, err, http.StatusForbidden)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseBody)
}

// SNS一覧を取得
func getAllSnsListHandler(w http.ResponseWriter, r *http.Request) {
	snsList, err := models.GetAllSnsList()
	if err != nil {
		ErrorHandler(w, err, http.StatusForbidden)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(snsList); err != nil {
		ErrorHandler(w, err, http.StatusForbidden)
		return
	}
}

// SNSの新規作成
func createSnsHandler(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		ErrorHandler(w, err, http.StatusForbidden)
		return
	}

	var sns models.Sns
	if err := json.Unmarshal(reqBody, &sns); err != nil {
		ErrorHandler(w, err, http.StatusForbidden)
		return
	}
	// modelの呼び出し
	err = models.InsertSns(&sns)
	if err != nil {
		ErrorHandler(w, err, http.StatusForbidden)
		return
	}
	responseBody, err := json.Marshal(sns)
	if err != nil {
		ErrorHandler(w, err, http.StatusForbidden)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseBody)
}
