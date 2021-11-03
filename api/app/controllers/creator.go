package controllers

import (
	"api/app/models"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// 製作者詳細を取得
func getCreatorHandler(w http.ResponseWriter, r *http.Request) {
	creator := models.GetCreator()
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(creator); err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("error: %s", err), http.StatusForbidden)
		return
	}
}

// 製作者詳細を更新
func updateCreatorHandler(w http.ResponseWriter, r *http.Request) {
	creator := models.GetCreator()
	// まずは製作者情報を更新

	// 終わったらロゴファイルを更新

	file, handler, err := r.FormFile("file")
	if file == nil {
		return
	}
	if err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("error: %s", err), http.StatusForbidden)

		return
	}
	defer file.Close()

	uuid, err := models.GenerateUuid()
	if err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("error: %s", err), http.StatusForbidden)

		return
	}
	savedirectory := fmt.Sprintf("img/%s/%s", uuid[0:1], uuid[1:2])
	// 保存用のディレクトリがない場合は作成する
	if err := os.MkdirAll(savedirectory, 0777); err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("error: %s", err), http.StatusForbidden)

		return
	}
	// fileのMIMETypeを取得
	mimeType := handler.Header["Content-Type"][0]
	savePath := savedirectory + "/" + uuid + typeToExtention[mimeType]
	f, err := os.OpenFile(savePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("error: %s", err), http.StatusForbidden)

		return
	}
	defer f.Close()
	io.Copy(f, file)

	// commitする

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(creator); err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("error: %s", err), http.StatusForbidden)
		return
	}
}
