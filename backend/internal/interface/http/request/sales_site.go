package request

type CreateSalesSiteRequest struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type UpdateSalesSiteRequest struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
