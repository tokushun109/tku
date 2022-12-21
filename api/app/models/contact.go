package models

type Contact struct {
	DefaultModel
	Name        string  `json:"name" validate:"min=1,max=20"`
	Company     *string `json:"company" validate:"min=1,max=20"`
	PhoneNumber *string `json:"phoneNumber" validate:"max=20"`
	Email       string  `json:"email" validate:"min=1,max=50,email"`
	Content     string  `json:"content"`
	IsRead      bool    `json:"isRead"`
}

func InsertContact(contact *Contact) (err error) {
	err = Db.Create(&contact).Error
	return err
}
