package models

import "time"

type Contact struct {
	DefaultModel
	ID          *uint   `gorm:"primary_key" json:"id"`
	Name        string  `json:"name" validate:"min=1,max=20"`
	Company     *string `json:"company" validate:"max=20"`
	PhoneNumber *string `json:"phoneNumber" validate:"max=20"`
	Email       string  `json:"email" validate:"min=1,max=50,email"`

	Content string `json:"content"`
	// レスポンスで返せるように上書き
	CreatedAt time.Time `json:"createdAt"`
}

type ContactList []Contact

func GetAllContactList() (contactList ContactList) {
	db := GetDBConnection()
	db.Order("created_at desc").Find(&contactList)

	return contactList
}

func InsertContact(contact *Contact) (err error) {
	db := GetDBConnection()
	err = db.Create(&contact).Error
	return err
}
