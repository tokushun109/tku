package models

type User struct {
	DefaultModel
	Uuid     string `json:"uuid"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Users []User

func GetAllUsers() (users Users) {
	Db.Find(&users)
	return users
}
