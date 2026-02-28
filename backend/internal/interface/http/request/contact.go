package request

type CreateContactRequest struct {
	Name        string  `json:"name"`
	Company     *string `json:"company"`
	PhoneNumber *string `json:"phoneNumber"`
	Email       string  `json:"email"`
	Content     string  `json:"content"`
}
