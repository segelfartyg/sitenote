package session

import (
	"net/http"
	"time"

	"github.com/google/uuid"
)

var sessions = map[string]Session{}

type Session struct {
	UserId string
	Expiry time.Time
}

func Setup() {

}

func Create(userId string) (Session, string) {

	sessionToken := uuid.NewString()
	expiryTime := time.Now().Add(120 * time.Second)

	sessions[sessionToken] = Session{
		UserId: userId,
		Expiry: expiryTime,
	}

	return sessions[sessionToken], sessionToken
}

func Validate(sessionToken string, w http.ResponseWriter) (userId string) {

	userSession, exists := sessions[sessionToken]

	if !exists {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if userSession.isExpired() {
		delete(sessions, sessionToken)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	return userSession.UserId

}

// CHECKING IF THE SESSION HAS EXPIRED
func (s Session) isExpired() bool {
	return s.Expiry.Before(time.Now())
}
