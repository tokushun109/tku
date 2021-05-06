package models

type User struct {
	DefaultModel
	Name     string `json:"name"`
	Email    string `json:"email"`
	PassWord string `json:"password"`
}
