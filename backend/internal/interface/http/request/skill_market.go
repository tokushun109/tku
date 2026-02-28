package request

type CreateSkillMarketRequest struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type UpdateSkillMarketRequest struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
