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
	if err := db.Begin().Error; err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("error: %s", err), http.StatusForbidden)
		return
	}
	db.Commit()
	// responseBodyで処理の成功を返す
	responseBody := getSuccessResponse()
	w.Header().Set("Content-Type", "application/json")
	w.Write(responseBody)
}
