package request

type LoginUserRequest struct {
	Email    string `json:"Email"`
	Password string `json:"Password"`
}
