package api

import (
	"net/http"
	"time"
)

// Rest logging middleware
func (a *API) loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func(begin time.Time) {
			a.logger.Log("[info]", "API Rest", "Method", r.RequestURI, "took", time.Since(begin))
		}(time.Now())
		next.ServeHTTP(w, r)
	})
}
