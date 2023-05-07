package consts

import (
	"fmt"
)

func RequestHeadersWithAuth(token string) map[string]string {
	authHeaders := map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", token),
	}
	// merge authHeaders and RequestHeaders
	for k, v := range RequestHeaders() {
		authHeaders[k] = v
	}

	return authHeaders
}

func RequestHeaders() map[string]string {
	return map[string]string{
		"Content-Type": "application/json",
	}
}

func BuildQueryURL(path string, queryParams map[string]string) string {
	queryURL := path
	if len(queryParams) > 0 {
		queryURL += "?"
		for k, v := range queryParams {
			queryURL += fmt.Sprintf("%s=%s&", k, v)
		}
		queryURL = queryURL[:len(queryURL)-1]
	}

	return queryURL
}
