package controllers

import (
	"api/app/models"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type loginForm struct {
	Email    string
	Password string
}

// ユーザー一覧を取得
func getAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	_, err := sessionCheck(r)
	if err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("error: %s", err), http.StatusForbidden)
		return
	}

	users := models.GetAllUsers()
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(users); err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("error: %s", err), http.StatusForbidden)
		return
	}
}

// ログインしているユーザーを取得
func getLoginUserHandler(w http.ResponseWriter, r *http.Request) {
	session, err := sessionCheck(r)
	if err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("error: %s", err), http.StatusForbidden)
		return
	}

	user := session.GetUserBySession()
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(user); err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("error: %s", err), http.StatusForbidden)
		return
	}
}

// ユーザーのログイン(セッションの作成)
func loginHandler(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("error: %s", err), http.StatusForbidden)
		return
	}

	var loginForm loginForm
	if err := json.Unmarshal(reqBody, &loginForm); err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("error: %s", err), http.StatusForbidden)
		return
	}

	user := models.GeUserByEmail(loginForm.Email)
	if user.Password == models.Encrypt(loginForm.Password) {
		// すでにuserに対応するsessionが作成されている場合は一度削除する
		if session := user.GetSessionByUser(); session.ID != nil {
			session.DeleteSession()
		}

		session, err := user.CreateSession()
		if err != nil {
			log.Println(err)
			http.Error(w, fmt.Sprintf("error: %s", err), http.StatusForbidden)
			return
		}

		responseBody, err := json.Marshal(session)
		if err != nil {
			log.Println(err)
			http.Error(w, fmt.Sprintf("error: %s", err), http.StatusForbidden)
			return

		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(responseBody)
	} else {
		err = errors.New("the password is incorrect")
		log.Println(err)
		http.Error(w, fmt.Sprintf("error: %s", err), http.StatusUnauthorized)
		return
	}
}

// ユーザーのログアウト(セッションの削除)
func logoutHandler(w http.ResponseWriter, r *http.Request) {
	session, err := sessionCheck(r)
	if err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("error: %s", err), http.StatusForbidden)
		return
	}

	if err := session.DeleteSession(); err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("error: %s", err), http.StatusForbidden)
		return
	}

}
