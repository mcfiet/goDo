package middleware

import (
	"net/http"
	"strconv"
)

func UseIsAuthenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if isAuthenticated(r) {
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, strconv.Itoa(http.StatusUnauthorized)+" - "+http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}
		return
	})
}

func isAuthenticated(r *http.Request) bool {
	// TODO: Implement isAuthenticated
	return true
}
