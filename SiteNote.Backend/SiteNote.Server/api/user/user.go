package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"notelad.com/server/db"
	"notelad.com/server/utils/cookie"
	"notelad.com/server/utils/session"
)

type UserCreateFindingRequest struct {
	Name    string `json:"name"`
	Link    string `json:"link"`
	Content string `json:"content"`
}

func UserCreateFinding(w http.ResponseWriter, r *http.Request) {

	cookie := cookie.Validate(w, r)
	sessionToken := cookie.Value
	userId := session.Validate(sessionToken, w)

	var req UserCreateFindingRequest

	e := json.NewDecoder(r.Body).Decode(&req)

	if e != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	newFindingId := uuid.NewString()

	db.CreateFinding(newFindingId, req.Name, req.Link, userId, req.Content)

	w.Write([]byte(fmt.Sprintf(newFindingId)))
}

func UserGetFindings(w http.ResponseWriter, r *http.Request) {
	cookie := cookie.Validate(w, r)
	sessionToken := cookie.Value
	userId := session.Validate(sessionToken, w)

	findings := db.GetUserFindings(userId)
	json.NewEncoder(w).Encode(findings)
}
