package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"cloud.google.com/go/auth/credentials/idtoken"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Audience string = "1010731658636-aeejci8n3gctj78iqdehtti3qfqpn568.apps.googleusercontent.com"

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

	userCollection = DB.Database("notelad").Collection("users")
	findingCollection = DB.Database("notelad").Collection("findings")

	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal(err)
	}

	if CheckIfUserExist("admin").UserId == "" {
		CreateUser("admin", "admin@notelad.com")
	}

	if CheckIfFindingExist("1").FindingId == "" {
		CreateFinding("1", "First Finding In NoteLad", "https://splodo.com", "admin")
	}

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

	sub := payload.Subject

	if CheckIfUserExist(sub).UserId == "" {
		CreateUser(sub, "no_email")
	}

	fmt.Print(sub)
	fmt.Print(payload.Claims)
	fmt.Print(payload.Claims)
}

func CheckIfUserExist(userId string) User {
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

func CreateUser(userId string, email string) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	res, err := userCollection.InsertOne(ctx, bson.D{{Key: "userId", Value: userId}, {Key: "email", Value: email}, {Key: "created", Value: time.Now().UTC().UnixMilli()}})

	if err != nil {
		log.Fatal(err)
	}

	id := res.InsertedID
	fmt.Print(id)
}

func CreateFinding(findingId string, name string, link string, userId string) {

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	res, err := findingCollection.InsertOne(ctx, bson.D{{Key: "findingId", Value: findingId}, {Key: "name", Value: name}, {Key: "link", Value: link}, {Key: "userId", Value: userId}, {Key: "created", Value: time.Now().UTC().UnixMilli()}})

	if err != nil {
		log.Fatal(err)
	}

	id := res.InsertedID
	fmt.Print(id)
}

func CheckIfFindingExist(findingId string) Finding {
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
