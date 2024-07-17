package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

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
	Content   string    `bson:"content"`
	Domain    string    `bson:"domain"`
	Created   time.Time `bson:"created"`
}

func Setup() {

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
	if CheckIfUserExist("admin").UserId == "" {
		CreateUser("admin", "admin@notelad.com")
	}

	if CheckIfFindingExist("1").FindingId == "" {
		CreateFinding("1", "First Finding In NoteLad", "https://splodo.com", "admin", "Remember to check this awesome website out", "splodo.com")
	}

}

// CHECKING IN DB FOR USER, RETURNING THE USER OBJECT IF PRESENT
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

// CREATING USER IN DATABASE
func CreateUser(userId string, email string) string {
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

// CHECKING IF FINDING EXISTS IN DB
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

// CREATING A FINDING IN DB
func CreateFinding(findingId, name, link, userId string, content string, domain string) {

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	res, err := findingCollection.InsertOne(ctx, bson.D{{Key: "findingId", Value: findingId}, {Key: "name", Value: name}, {Key: "link", Value: link}, {Key: "domain", Value: domain}, {Key: "userId", Value: userId}, {Key: "content", Value: content}, {Key: "created", Value: time.Now().UTC().UnixMilli()}})

	if err != nil {
		log.Fatal(err)
	}

	id := res.InsertedID
	fmt.Print(id)
}

func GetUserFindings(userId string) []Finding {
	var result []Finding

	filter := bson.D{{Key: "userId", Value: userId}}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cursor, err := findingCollection.Find(ctx, filter)

	if err == mongo.ErrNoDocuments {
		fmt.Println("records does not exist")
	} else if err != nil {
		log.Fatal(err)
	}

	if err = cursor.All(ctx, &result); err != nil {
		panic(err)
	}

	return result
}

func GetUserFindingsFromDomain(userId string, domain string) []Finding {
	var result []Finding
	fmt.Println(userId)
	fmt.Println(domain)
	filter := bson.D{{Key: "userId", Value: userId}, {Key: "domain", Value: domain}}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cursor, err := findingCollection.Find(ctx, filter)

	if err == mongo.ErrNoDocuments {
		fmt.Println("records does not exist")
	} else if err != nil {
		log.Fatal(err)
	}

	if err = cursor.All(ctx, &result); err != nil {
		panic(err)
	}

	return result
}

func GetUserFinding(userId string, findingId string) Finding {
	var result Finding

	filter := bson.D{{Key: "userId", Value: userId}, {Key: "findingId", Value: findingId}}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	err := findingCollection.FindOne(ctx, filter).Decode(&result)

	if err != nil {
		panic(err)
	}

	return result
}

func DeleteUserFindings(userId string, findingId string) string {

	filter := bson.D{{Key: "userId", Value: userId}, {Key: "findingId", Value: findingId}}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	_, err := findingCollection.DeleteOne(ctx, filter)

	if err == mongo.ErrNoDocuments {
		fmt.Println("records does not exist")
	} else if err != nil {
		log.Fatal(err)
	}

	return findingId
}
