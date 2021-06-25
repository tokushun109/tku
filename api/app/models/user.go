package models

import (
	"crypto/sha1"
	"fmt"
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

func (session *Session) GetUserBySession() (user User, err error) {
	err = Db.First(&user, "id = ?", session.UserId).Error
	return user, err

}

func GetAllUsers() (users Users, err error) {
	err = Db.Find(&users).Error
	return users, err
}

func GeUserByEmail(email string) (user User, err error) {
	err = Db.First(&user, "email = ?", email).Error
	return user, err
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

func (user *User) GetSessionByUser() (session Session, err error) {
	err = Db.First(&session, "user_id = ?", user.ID).Error
	return session, err

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

func GetSession(uuid string) (session Session, err error) {
	err = Db.First(&session, "uuid = ?", uuid).Error
	return session, err
}

func (session *Session) IsValidSession() (valid bool, err error) {
	valid = false
	err = Db.First(&session, "uuid = ?", session.Uuid).Error
	if err != nil {
		return valid, err
	}
	if session != nil {
		valid = true
	}
	return valid, err
}

func (session *Session) DeleteSession() (err error) {
	err = Db.Delete(&session).Error
	return err
}
