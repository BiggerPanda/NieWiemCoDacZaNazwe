package models

import (
	"github.com/kamva/mgm/v3"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	mgm.DefaultModel `bson:",inline"`
	Username         string `json:"username" bson:"username"`
	Password         string `json:"password" bson:"password"`
	Email            string `json:"email" bson:"email"`
}

func CreateNewUser(username string, password string, email string) *User {
	pw, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	if err != nil {
		panic(err)
	}
	return &User{
		Username: username,
		Password: string(pw),
		Email:    email,
	}
}
