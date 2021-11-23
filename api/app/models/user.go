package models

import (
	"crypto/sha1"
	"fmt"
	"time"
)

type User struct {
	DefaultModel
	Uuid     string `json:"uuid"`
	Name     string `json:"name" validate:"min=1,max=20"`
	Email    string `json:"email" validate:"email"`
	Password string `json:"-" validate:"min=1,max=20"`
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
	Db.Limit(1).Find(&user, "id = ?", session.UserId)
	return user

}

func GetAllUsers() (users Users) {
	Db.Find(&users)
	return users
}

func GeUserByEmail(email string) (user User) {
	Db.Limit(1).Find(&user, "email = ?", email)
	return user
}

// is_adminは管理ユーザーならtrue、一般ユーザーならfalse
func InsertUser(user *User, is_admin bool) (err error) {
	user.IsAdmin = is_admin
	// uuidの設定
	uuid, err := GenerateUuid()
	if err != nil {
		return err
	}
	user.Uuid = uuid
	// パスワードの変換を行う
	user.Password = Encrypt(user.Password)
	err = Db.Create(&user).Error
	return err
}

func (user *User) GetSessionByUser() (session Session) {
	Db.Limit(1).Find(&session, "user_id = ?", user.ID)
	return session

}

// Sessionデータを作成する
func (user *User) CreateSession() (session Session, err error) {
	session = Session{}
	uuid, err := GenerateUuid()
	if err != nil {
		return session, err
	}

	session.Uuid = uuid
	session.UserId = user.ID
	session.CreatedAt = time.Now()

	err = Db.Create(&session).Error

	return session, err
}

func GetSession(uuid string) (session Session) {
	Db.Limit(1).Find(&session, "uuid = ?", uuid)
	return session
}

func (session *Session) IsValidSession() (valid bool) {
	valid = false
	Db.Limit(1).Find(&session, "uuid = ?", session.Uuid)
	if session != nil {
		valid = true
	}
	return valid
}

func (session *Session) DeleteSession() (err error) {
	err = Db.Delete(&session).Error
	return err
}
