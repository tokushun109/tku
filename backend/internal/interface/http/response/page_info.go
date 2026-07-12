package response

type CursorPageInfoResponse struct {
	HasMore    bool   `json:"hasMore"`
	NextCursor string `json:"nextCursor"`
}

type OffsetPageInfoResponse struct {
	Page       int `json:"page"`
	Limit      int `json:"limit"`
	Total      int `json:"total"`
	TotalPages int `json:"totalPages"`
}
