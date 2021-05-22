package controllers

import (
	"api/app/models"
	"encoding/json"
	"log"
	"net/http"
)

// 商品一覧を取得
func getAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	users := models.GetAllUsers()
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(users); err != nil {
		log.Fatalln(err)
	}
}
