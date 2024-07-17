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

type UserGetFindingRequest struct {
	FindingId string `json:"findingId"`
}

type UserDeleteFindingRequest struct {
	FindingId string `json:"findingId"`
}

func UserCreateFinding(w http.ResponseWriter, r *http.Request) {

	sessionToken, err := cookie.RetrieveSessionToken(w, r)

	if err != nil {
		return
	}
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
	sessionToken, err := cookie.RetrieveSessionToken(w, r)

	if err != nil {
		return
	}

	userId := session.Validate(sessionToken, w)

	findings := db.GetUserFindings(userId)
	json.NewEncoder(w).Encode(findings)
}

func UserGetFinding(w http.ResponseWriter, r *http.Request) {
	sessionToken, err := cookie.RetrieveSessionToken(w, r)

	if err != nil {
		return
	}

	var req UserGetFindingRequest

	e := json.NewDecoder(r.Body).Decode(&req)

	if e != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	userId := session.Validate(sessionToken, w)

	finding := db.GetUserFinding(userId, req.FindingId)
	json.NewEncoder(w).Encode(finding)
}

func UserDeleteFinding(w http.ResponseWriter, r *http.Request) {
	sessionToken, err := cookie.RetrieveSessionToken(w, r)

	if err != nil {
		return
	}

	userId := session.Validate(sessionToken, w)

	var req UserDeleteFindingRequest

	e := json.NewDecoder(r.Body).Decode(&req)

	if e != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	deletedUserFinding := db.DeleteUserFindings(userId, req.FindingId)
	json.NewEncoder(w).Encode(deletedUserFinding)
}
