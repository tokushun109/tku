package controllers

import (
	"api/app/models"
	"encoding/json"
	"net/http"
)

// 製作者詳細を取得
func getCreatorHandler(w http.ResponseWriter, r *http.Request) {
	creator, err := models.GetCreator()
	if err != nil {
		ErrorHandler(w, err, http.StatusForbidden)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(creator); err != nil {
		ErrorHandler(w, err, http.StatusForbidden)
		return
	}
}
