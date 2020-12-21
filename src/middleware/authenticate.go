package middleware

import (
	"helper"
	"net/http"
	"strings"
)

// Authenticate checks http headers and either runs callback function or 401 status code if auth fails.
func Authenticate(settings helper.CommanderSettings, w http.ResponseWriter, r *http.Request, callback func(w http.ResponseWriter, r *http.Request)) {
	authHeader := r.Header.Get("authorization")
	apiKeyHeader := strings.Split(authHeader, " ")

	if (len(apiKeyHeader) == 2 && apiKeyHeader[1] == settings.APIKey) || settings.NoAuth {
		callback(w, r)
		return
	}

	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte("401 - Unauthorized"))
}
