package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient *mongo.Client
var rdb *redis.Client
var ctx = context.TODO()
var mongodb *mongo.Database

func main() {
	RedisConnect()
	initMongoConnection()
	handleFuncs()

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func handleFuncs() {
	router := mux.NewRouter()
	fileServer := http.FileServer(http.Dir("./static"))
	router.Handle("/", fileServer)
	router.HandleFunc("/CreateUser", HandleCreateUser).Methods("POST")
	http.Handle("/", router)
}

func RedisConnect() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	fmt.Println(rdb.ClientGetName(ctx))
	fmt.Println("RedisConnected")
}

func initMongoConnection() {
	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoConnected")

	err = mongoClient.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	mongodb = mongoClient.Database("TestDatabase")

	CreateUser(&User{ID: primitive.NewObjectID(),
		Username: "test",
		Password: "test",
		Email:    "test",
		Code:     "test"})
}
