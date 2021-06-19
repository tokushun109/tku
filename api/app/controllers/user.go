package controllers

import (
	"api/app/models"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type loginForm struct {
	Email    string
	Password string
}

// 商品一覧を取得
func getAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	users := models.GetAllUsers()
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(users); err != nil {
		log.Fatalln(err)
	}
}

// ログインしているユーザーを取得
func getLoginUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uuid := vars["session_uuid"]
	// ログインしているかの確認
	session, err := sessionCheck(uuid)
	if err == nil {
		user := session.GetUserBySession()
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(user); err != nil {
			log.Fatalln(err)
		}
	}
}

// ユーザーのログインとセッションの作成
func loginHandler(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	var loginForm loginForm
	if err := json.Unmarshal(reqBody, &loginForm); err != nil {
		log.Fatal(err)
	}

	user := models.GeUserByEmail(loginForm.Email)
	if user.ID == nil {
		log.Fatalln("ユーザーが登録されていません")
	}

	if user.Password == models.Encrypt(loginForm.Password) {
		// すでにuserに対応するsessionが作成されている場合は一度削除する
		if session := user.GetSessionByUser(); session.ID != nil {
			session.DeleteSession()
		}

		session := user.CreateSession()
		responseBody, err := json.Marshal(session)
		if err != nil {
			log.Fatal(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(responseBody)
	} else {
		log.Fatalln("パスワードが間違っています")
	}
}
