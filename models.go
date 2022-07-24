package main

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID
	Username string
	Email    string
	Password string
	Code     string
}
