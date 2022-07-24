package main

import (
	"encoding/json"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func HandleErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user User
	err := decoder.Decode(&user)
	if err != nil {
		// bad JSON or unrecognized json field
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if user.Code == "" {
		http.Error(w, "missing field 'test' from JSON object", http.StatusBadRequest)
		return
	}

	// optional extra check
	if decoder.More() {
		http.Error(w, "extraneous data after JSON object", http.StatusBadRequest)
		return
	}
	log.Println(user.Username)
	CreateUser(&user)

}

func HashAndSalt(pass []byte) string {
	hashed, err := bcrypt.GenerateFromPassword(pass, bcrypt.MinCost)
	HandleErr(err)

	return string(hashed)
}

func CreateUser(user *User) error {
	_, err := mongodb.Collection("Users").InsertOne(ctx, user)
	rdb.Set(ctx, user.ID.String(), user.Code, 0)
	return err
}
