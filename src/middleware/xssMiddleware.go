package middleware

import (
	"context"
	"html"
	"net/http"
)

type contextKey string

const (
	cleanedInputKey contextKey = "cleanedInput"
)

func XssMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		inputValue := r.URL.Query().Get("input")
		cleanedInput := html.EscapeString(inputValue)
		ctx := context.WithValue(r.Context(), cleanedInputKey, cleanedInput)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

func GetCleanedInput(r *http.Request) string {
	if value, ok := r.Context().Value(cleanedInputKey).(string); ok {
		return value
	}
	return ""
}
