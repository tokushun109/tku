package request

type CreateSnsRequest struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type UpdateSnsRequest struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
