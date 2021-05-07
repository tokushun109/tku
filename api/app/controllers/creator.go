package controllers

import (
	"api/app/models"
	"encoding/json"
	"log"
	"net/http"
)

// 製作者詳細を取得
func getCreatorHandler(w http.ResponseWriter, r *http.Request) {
	creator := models.GetCreator()
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(creator)
	if err != nil {
		log.Fatalln(err)
	}
}
