package controllers

import (
	"api/app/models"
	"fmt"
	"log"
	"net/http"
)

// サーバーのヘルスチェック情報を取得
func getHealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	// DBの接続チェック
	db := models.GetDBConnection()
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("error: %s", err), http.StatusForbidden)
		return
	}
	tx.Commit()
	// responseBodyで処理の成功を返す
	responseBody := getSuccessResponse()
	w.Header().Set("Content-Type", "application/json")
	w.Write(responseBody)
}
