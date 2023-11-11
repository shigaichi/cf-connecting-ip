package cfconnectingip

import "net/http"

func SetRemoteAddr(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if cfIP := r.Header.Get("CF-Connecting-IP"); cfIP != "" {
			r.RemoteAddr = cfIP
		}
		next.ServeHTTP(w, r)
	})
}
