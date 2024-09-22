package gocardless

import (
	"fmt"
	"strings"
	"time"
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

type TimeWithTimeZoneInfo struct {
	time.Time
}

const timeWithTimeZoneInfo = "2006-01-02T15:04:05.999999"

func (ct *TimeWithTimeZoneInfo) UnmarshalJSON(b []byte) error {
	str := strings.Trim(string(b), `"`)
	t, err := time.Parse(timeWithTimeZoneInfo, str)
	if err != nil {
		return fmt.Errorf("failed to parse time: %w", err)
	}
	ct.Time = t
	return nil
}

type TimeWithTimeZoneInfoZ struct {
	time.Time
}

const timeWithTimeZoneInfoZ = "2006-01-02T15:04:05.999999Z"

func (ct *TimeWithTimeZoneInfoZ) UnmarshalJSON(b []byte) error {
	str := strings.Trim(string(b), `"`)
	t, err := time.Parse(timeWithTimeZoneInfoZ, str)
	if err != nil {
		return fmt.Errorf("failed to parse time: %w", err)
	}
	ct.Time = t
	return nil
}
