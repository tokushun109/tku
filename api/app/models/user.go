package models

import (
	"crypto/sha1"
	"fmt"
	"log"
	"time"
)

type User struct {
	DefaultModel
	Uuid     string `json:"uuid"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
	IsAdmin  bool   `json:"isAdmin"`
}

type Users []User

// TODO emailのfieldをなくす
type Session struct {
	ID        *uint     `gorm:"primary_key" json:"-"`
	Uuid      string    `json:"uuid"`
	UserId    *uint     `json:"-"`
	CreatedAt time.Time `json:"-"`
}

func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return cryptext
}

func (session *Session) GetUserBySession() (user User) {
	Db.First(&user, "id = ?", session.UserId)
	return user

}

func GetAllUsers() (users Users) {
	Db.Find(&users)
	return users
}

func GeUserByEmail(email string) (user User) {
	Db.First(&user, "email = ?", email)
	return user
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

func (user *User) GetSessionByUser() (session Session) {
	Db.First(&session, "user_id = ?", user.ID)
	return session

}

// Sessionデータを作成する
func (user *User) CreateSession() (session Session) {
	session = Session{}
	uuid, err := GenerateUuid()
	if err != nil {
		log.Fatal(err)
	}

	session.Uuid = uuid
	session.UserId = user.ID
	session.CreatedAt = time.Now()

	Db.NewRecord(session)
	Db.Create(&session)

	return session
}

func GetSession(uuid string) (session Session) {
	Db.First(&session, "uuid = ?", uuid)
	return session
}

func (session *Session) IsValidSession() (valid bool) {
	valid = false
	Db.First(&session, "uuid = ?", session.Uuid)
	if session != nil {
		valid = true
	}
	return valid
}

func (session *Session) DeleteSession() {
	Db.Delete(&session)
}
