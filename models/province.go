package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

//Province object
type Province struct {
	ID         bson.ObjectId `json:"id" bson:"_id"`
	Name       string        `json:"name" bson:"name"`
	Intro      string        `json:"intro" bson:"intro"`
	Content    string        `json:"content" bson:"content"`
	CoverImage string        `json:"cover_image" bson:"cover_image"`
	CreatedOn  time.Time     `json:"created_on" bson:"created_on"`
	UpdatedOn  time.Time     `json:"updated_on" bson:"updated_on"`
}

//Provinces is a slices of Province
type Provinces []Province
