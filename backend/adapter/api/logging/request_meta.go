package logging

import "net/http"

func getRequestMeta(r *http.Request) (method, path string) {
	if r != nil {
		method = r.Method
		if r.URL != nil {
			path = r.URL.String()
		}
	}
	return
}
