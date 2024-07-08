package auth

import (
	"errors"
	"net/http"
	"strings"
)

// Auth -> api-key {insert API key}
func GetApiKey(headers http.Header) (string, error) {

	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no authentication info found")
	}

	vals := strings.Split(val, " ")

	if len(vals) != 2 {
		return "", errors.New("malformed auth header")
	}
	if vals[0] != "api-key" {
		return "", errors.New("malformed auth header")
	}
	if len(vals[1]) != 64 {
		return "", errors.New("invalid api-key")
	}

	return vals[1], nil
}
