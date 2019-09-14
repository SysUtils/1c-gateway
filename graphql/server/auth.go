package server

import (
	"net/http"
)

type ITokenManager interface {
	Get(login, password string) (string, error)
	Check(token string) bool
}

func Secure(handler http.Handler, manager ITokenManager) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		valid := manager.Check(token)
		if !valid {
			w.WriteHeader(401)
			return
		}

		handler.ServeHTTP(w, r)
	})
}
