package request

type CreateSalesSiteRequest struct {
	Name string `json:"name"`
	URL  string `json:"url"`
	Icon string `json:"icon"`
}

type UpdateSalesSiteRequest struct {
	Name string `json:"name"`
	URL  string `json:"url"`
	Icon string `json:"icon"`
}
