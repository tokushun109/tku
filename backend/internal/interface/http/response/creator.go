package response

type CreatorResponse struct {
	Name         string `json:"name"`
	Introduction string `json:"introduction"`
	MimeType     string `json:"mimeType"`
	Logo         string `json:"logo"`
	APIPath      string `json:"apiPath"`
}
