package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"cloud.google.com/go/auth/credentials/idtoken"
	"github.com/gorilla/mux"
)

var Audience string = "1010731658636-aeejci8n3gctj78iqdehtti3qfqpn568.apps.googleusercontent.com"

type LoginRequest struct {
	IdToken string `json:"id_token"`
}

func main() {

	fmt.Println("Starting server")

	router := mux.NewRouter()
	router.HandleFunc("/login", Login)

	http.ListenAndServe(":9000", corsMiddleware(router))
}

func Login(w http.ResponseWriter, r *http.Request) {

	var req LoginRequest

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	fmt.Print(req.IdToken)
	ValidateGoogleIdToken(req.IdToken)
	fmt.Fprint(w, "GOT THIS: "+req.IdToken)

}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Executing middleware", r.Method)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers:", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func ValidateGoogleIdToken(token string) {
	payload, err := idtoken.Validate(context.Background(), token, Audience)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(payload.Claims)
}
