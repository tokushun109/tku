package response

type CursorPageInfoResponse struct {
	HasMore    bool   `json:"hasMore"`
	NextCursor string `json:"nextCursor"`
}
