package consts

import (
	"fmt"
	"os"
	"testing"
)

func GetSecrets(t *testing.T) (string, string) {
	t.Helper()

	secretID := os.Getenv("NORDIGEN_SECRET_ID")
	secretKey := os.Getenv("NORDIGEN_SECRET_KEY")
	if secretID == "" || secretKey == "" {
		fmt.Println("NORDIGEN_SECRET_ID or NORDIGEN_SECRET_KEY is not set")
	}

	return secretID, secretKey
}

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
