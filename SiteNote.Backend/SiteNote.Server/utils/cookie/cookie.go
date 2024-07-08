package cookie

import "net/http"

func Validate(w http.ResponseWriter, r *http.Request) *http.Cookie {
	c, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
		}
		w.WriteHeader(http.StatusBadRequest)
	}
	return c
}
