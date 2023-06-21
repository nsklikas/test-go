package middlewares

import "net/http"

type AuthMiddleware struct {
	AuthToken string
}

func (s AuthMiddleware) Run(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token != s.AuthToken {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}
