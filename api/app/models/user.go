package models

type User struct {
	DefaultModel
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Users []User

func GetAllUsers() (users Users) {
	Db.Find(&users)
	return users
}
