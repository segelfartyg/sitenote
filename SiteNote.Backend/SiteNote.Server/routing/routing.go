package routing

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"notelad.com/server/api/auth"
	"notelad.com/server/api/user"
	"notelad.com/server/consts"
)

func Setup() {

	// STARTING HTTP SERVER
	fmt.Println("Starting server")

	// ROUTE HANDLING
	router := mux.NewRouter()

	router.HandleFunc("/login", auth.Login)
	router.HandleFunc("/getUser", auth.GetUserId)
	router.HandleFunc("/finding/create", user.UserCreateFinding)
	router.HandleFunc("/finding/user", user.UserGetFinding)
	router.HandleFunc("/finding/user/all", user.UserGetFindings)
	router.HandleFunc("/finding/user/domain/all", user.UserGetFindingsFromDomain)
	router.HandleFunc("/finding/user/delete", user.UserDeleteFinding)

	// STARTING SERVER WITH ROUTER + CORS CONFIG
	http.ListenAndServe(consts.Port, corsMiddleware(router))
}

// CORS MIDDLEWARE FOR ENABLING CLIENT ACCESS TO API
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		log.Println("Executing middleware", r.Method)

		switch origin := r.Header.Get("Origin"); origin {
		case "chrome-extension://cpediolkjjaolfdjmgkhaaglfgfgejld", "http://localhost:5173":
			w.Header().Set("Access-Control-Allow-Origin", origin)
		}

		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers:", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
