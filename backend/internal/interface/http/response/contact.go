package response

import "time"

type ContactResponse struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Company     *string   `json:"company"`
	PhoneNumber *string   `json:"phoneNumber"`
	Email       string    `json:"email"`
	Content     string    `json:"content"`
	CreatedAt   time.Time `json:"createdAt"`
}
