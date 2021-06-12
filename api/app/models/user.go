package models

import (
	"crypto/sha1"
	"fmt"
	"log"
)

type User struct {
	DefaultModel
	Uuid     string `json:"uuid"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	IsAdmin  bool   `json:"is_admin"`
}

type Users []User

func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return cryptext
}

func GetAllUsers() (users Users) {
	Db.Find(&users)
	return users
}

// is_adminは管理ユーザーならtrue、一般ユーザーならfalse
func InsertUser(user *User, is_admin bool) {
	user.IsAdmin = is_admin
	// uuidの設定
	uuid, err := GenerateUuid()
	if err != nil {
		log.Fatal(err)
	}
	user.Uuid = uuid
	// パスワードの変換を行う
	user.Password = Encrypt(user.Password)
	Db.NewRecord(user)
	Db.Create(&user)
}
