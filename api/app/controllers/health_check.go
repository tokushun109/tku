package controllers

import (
	"net/http"
)

// サーバーのヘルスチェック情報を取得
func getHealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	// responseBodyで処理の成功を返す
	responseBody := getSuccessResponse()
	w.Header().Set("Content-Type", "application/json")
	w.Write(responseBody)
}
