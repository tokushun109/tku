package request

type CreateSnsRequest struct {
	Name string `json:"name"`
	URL  string `json:"url"`
	Icon string `json:"icon"`
}

type UpdateSnsRequest struct {
	Name string `json:"name"`
	URL  string `json:"url"`
	Icon string `json:"icon"`
}
