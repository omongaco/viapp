package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

//User object
type User struct {
	ID          bson.ObjectId `json:"id" bson:"_id"`
	Firstname   string        `json:"name" bson:"name"`
	Lastname    string        `json:"intro" bson:"intro"`
	Username    string        `json:"content" bson:"content"`
	Email       string        `json:"email" bson:"email"`
	PasswordMD5 string        `json:"password_md5" bson:"password_md5"`
	Photo       string        `json:"cover_image" bson:"cover_image"`
	LastLogin   time.Time     `json:"last_login" bson:"last_login"`
	CreatedOn   time.Time     `json:"created_on" bson:"created_on"`
	UpdatedOn   time.Time     `json:"updated_on" bson:"updated_on"`
}

//Users is a slices of User
type Users []User
