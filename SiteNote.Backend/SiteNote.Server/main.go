package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"cloud.google.com/go/auth/credentials/idtoken"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Audience string = "1010731658636-aeejci8n3gctj78iqdehtti3qfqpn568.apps.googleusercontent.com"
var sessions = map[string]session{}

type session struct {
	userId string
	expiry time.Time
}

var (
	DB                *mongo.Client
	userCollection    *mongo.Collection
	findingCollection *mongo.Collection
)

type User struct {
	UserId  string    `bson:"userId"`
	Email   string    `bson:"email"`
	Created time.Time `bson:"created"`
}

type Finding struct {
	FindingId string    `bson:"findingId"`
	Name      string    `bson:"name"`
	Link      string    `bson:"link"`
	UserId    string    `bson:"userId"`
	Created   time.Time `bson:"created"`
}

type LoginRequest struct {
	IdToken string `json:"id_token"`
}

func main() {

	// SETTING UP DB CLIENT
	fmt.Println("Connecting to DB")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	DB, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	// SETTING UP DB COLLECTIONS
	userCollection = DB.Database("notelad").Collection("users")
	findingCollection = DB.Database("notelad").Collection("findings")

	if err != nil {
		log.Fatal(err)
	}

	// SETTING UP ROOT USER AND FINDING
	if checkIfUserExist("admin").UserId == "" {
		createUser("admin", "admin@notelad.com")
	}

	if checkIfFindingExist("1").FindingId == "" {
		createFinding("1", "First Finding In NoteLad", "https://splodo.com", "admin")
	}

	// STARTING HTTP SERVER
	fmt.Println("Starting server")

	// ROUTE HANDLING
	router := mux.NewRouter()

	router.HandleFunc("/login", login)
	router.HandleFunc("/getUser", getUserId)

	// STARTING SERVER WITH ROUTER + CORS CONFIG
	http.ListenAndServe(":9000", corsMiddleware(router))
}

// LOGIN, CHECKING FOR GOOGLE ID TOKEN, THEN AUTHENTICATING
func login(w http.ResponseWriter, r *http.Request) {

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

	if checkIfUserExist(googleUserId).UserId == "" {
		createUser(googleUserId, "no_email")
	}

	sessionToken := uuid.NewString()
	expiresAt := time.Now().Add(120 * time.Second)

	sessions[sessionToken] = session{
		userId: googleUserId,
		expiry: expiresAt,
	}
	fmt.Println("SESSIONS HALLÅ")
	fmt.Println(sessions)

	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		SameSite: http.SameSiteNoneMode,
		Secure:   true,
		HttpOnly: true,
		Value:    sessionToken,
		Expires:  expiresAt,
	})

	//fmt.Fprint(w, "GOT THIS: "+req.IdToken)

}

// CORS MIDDLEWARE FOR ENABLING CLIENT ACCESS TO API
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Executing middleware", r.Method)
		w.Header().Set("Access-Control-Allow-Origin", "chrome-extension://cpediolkjjaolfdjmgkhaaglfgfgejld")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers:", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

// VALIDATION FUNCTION FOR GOOGLE ID TOKENS, CHECKING AGAINST AUDIENCE AS WELL
func validateGoogleIdToken(token string) string {
	payload, err := idtoken.Validate(context.Background(), token, Audience)

	if err != nil {
		log.Fatal(err)
		return ""
	}

	sub := payload.Subject

	return sub
}

// CHECKING IN DB FOR USER, RETURNING THE USER OBJECT IF PRESENT
func checkIfUserExist(userId string) User {
	var result User

	filter := bson.D{{Key: "userId", Value: userId}}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	err := userCollection.FindOne(ctx, filter).Decode(&result)

	if err == mongo.ErrNoDocuments {
		// Do something when no record was found
		fmt.Println("record does not exist")
	} else if err != nil {
		log.Fatal(err)
	}

	return result
}

// CREATING USER IN DATABASE
func createUser(userId string, email string) string {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	res, err := userCollection.InsertOne(ctx, bson.D{{Key: "userId", Value: userId}, {Key: "email", Value: email}, {Key: "created", Value: time.Now().UTC().UnixMilli()}})

	if err != nil {
		log.Fatal(err)
	}

	id := res.InsertedID
	// LOGGING MONGODB ObjectID
	fmt.Print(id)
	return userId
}

// CREATING A FINDING IN DB
func createFinding(findingId, name, link, userId string) {

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	res, err := findingCollection.InsertOne(ctx, bson.D{{Key: "findingId", Value: findingId}, {Key: "name", Value: name}, {Key: "link", Value: link}, {Key: "userId", Value: userId}, {Key: "created", Value: time.Now().UTC().UnixMilli()}})

	if err != nil {
		log.Fatal(err)
	}

	id := res.InsertedID
	fmt.Print(id)
}

// CHECKING IF FINDING EXISTS IN DB
func checkIfFindingExist(findingId string) Finding {
	var result Finding

	filter := bson.D{{Key: "findingId", Value: findingId}}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	err := findingCollection.FindOne(ctx, filter).Decode(&result)

	if err == mongo.ErrNoDocuments {
		// Do something when no record was found
		fmt.Println("record does not exist")
	} else if err != nil {
		log.Fatal(err)
	}

	return result
}

// ROUTE HANDLER FOR GETTING USERID FROM USER SESSION
func getUserId(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
		}
		//w.WriteHeader(http.StatusBadRequest)
	}

	sessionToken := c.Value
	userSession, exists := sessions[sessionToken]
	if !exists {
		w.WriteHeader(http.StatusUnauthorized)
	}

	if userSession.isExpired() {
		delete(sessions, sessionToken)
		w.WriteHeader(http.StatusUnauthorized)
	}

	w.Write([]byte(fmt.Sprintf("USERID: %s", userSession.userId)))
}

// CHECKING IF THE SESSION HAS EXPIRED
func (s session) isExpired() bool {
	return s.expiry.Before(time.Now())
}
