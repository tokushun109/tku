package controllers

import (
	"api/app/models"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
	"gopkg.in/go-playground/validator.v9"
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

// ロゴ画像のパスからバイナリデータを返す
func getCreatorLogoBlobHandler(w http.ResponseWriter, r *http.Request) {
	creator := models.GetCreator()
	vars := mux.Vars(r)
	// リクエストされているロゴの名前を取得
	requestLogoFile := vars["logo_file"]
	// 保存されているロゴの名前を取得
	logoFile := strings.Split(creator.Logo, "/")[4]

	// リクエストされたロゴファイルと現状の製作者のロゴが異なる場合はエラーを返す
	if requestLogoFile != logoFile {
		err := errors.New("the request is invalid")
		log.Println(err)
		http.Error(w, fmt.Sprintf("error: %s", err), http.StatusBadRequest)
		return
	}

	file, err := os.Open(creator.Logo)
	if err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("error: %s", err), http.StatusForbidden)
		return
	}
	defer file.Close()

	binary, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("error: %s", err), http.StatusForbidden)
		return
	}
	w.Header().Set("Content-Type", creator.MimeType)
	w.Write(binary)
}

// 製作者詳細を更新
func updateCreatorHandler(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("error: %s", err), http.StatusForbidden)
		return
	}
	var creator models.Creator
	if err := json.Unmarshal(reqBody, &creator); err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("error: %s", err), http.StatusForbidden)
		return
	}

	// validationの確認
	validate := validator.New()
	if err := validate.Struct(creator); err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("error: %s", err), http.StatusBadRequest)
		return
	}

	if err = models.UpdateCreator(&creator); err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("error: %s", err), http.StatusBadRequest)
		return
	}

	// responseBodyで処理の成功を返す
	responseBody := getSuccessResponse()
	w.Header().Set("Content-Type", "application/json")
	w.Write(responseBody)
}

// 製作者ロゴを更新
func updateCreatorLogoHandler(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("logo")
	if err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("error: %s", err), http.StatusForbidden)
		return
	}
	uuid, err := models.GenerateUuid()
	if err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("error: %s", err), http.StatusForbidden)
		return
	}
	savedirectory := fmt.Sprintf("img/logo/%s/%s", uuid[0:1], uuid[1:2])
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

	// ファイルの情報をsqlに保存する
	var creator models.Creator
	// creatorのfieldを更新する
	creator.MimeType = mimeType
	creator.Logo = savePath
	// sqlにデータを作成する
	err = models.UpdateCreatorLogo(&creator)
	if err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("error: %s", err), http.StatusForbidden)
		return
	}

	// responseBodyで処理の成功を返す
	responseBody := getSuccessResponse()
	w.Header().Set("Content-Type", "application/json")
	w.Write(responseBody)
}
