package cookie

import "net/http"

func RetrieveSessionToken(w http.ResponseWriter, r *http.Request) (string, error) {
	c, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return "", err
		}
		w.WriteHeader(http.StatusBadRequest)
		return "", err
	}
	return c.Value, err
}
