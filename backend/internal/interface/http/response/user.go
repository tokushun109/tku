package response

type CurrentUserResponse struct {
	UUID    string `json:"uuid"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	IsAdmin bool   `json:"isAdmin"`
}

type LoginSessionResponse struct {
	UUID string `json:"uuid"`
}
