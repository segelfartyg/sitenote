package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"google.golang.org/api/idtoken"
	"notelad.com/server/consts"
	"notelad.com/server/db"
	"notelad.com/server/utils/cookie"
	"notelad.com/server/utils/session"
)

type LoginRequest struct {
	IdToken string `json:"id_token"`
}

func Login(w http.ResponseWriter, r *http.Request) {

	var req LoginRequest

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	fmt.Print(req.IdToken)
	var googleUserId = validateGoogleIdToken(req.IdToken)

	if googleUserId == "" {
		log.Fatal("Google Id Validation failed")

	}

	if db.CheckIfUserExist(googleUserId).UserId == "" {
		db.CreateUser(googleUserId, "no_email")
	}

	createdSession, sessionKey := session.Create(googleUserId)

	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		SameSite: http.SameSiteNoneMode,
		Secure:   true,
		HttpOnly: true,
		Value:    sessionKey,
		Expires:  createdSession.Expiry,
	})

	//RETURNING USERID
	fmt.Fprint(w, googleUserId)

}

// ROUTE HANDLER FOR GETTING USERID FROM USER SESSION
func GetUserId(w http.ResponseWriter, r *http.Request) {
	sessionToken, err := cookie.RetrieveSessionToken(w, r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	userId := session.Validate(sessionToken, w)
	w.Write([]byte(fmt.Sprintf(userId)))
}

// VALIDATION FUNCTION FOR GOOGLE ID TOKENS, CHECKING AGAINST AUDIENCE AS WELL
func validateGoogleIdToken(token string) string {
	payload, err := idtoken.Validate(context.Background(), token, consts.Audience)
	if err != nil {
		log.Fatal(err)
		return ""
	}
	sub := payload.Subject
	return sub
}
