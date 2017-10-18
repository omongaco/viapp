package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

//Activity object
type Activity struct {
	ID         bson.ObjectId `json:"id" bson:"_id"`
	Name       string        `json:"name" bson:"name"`
	Intro      string        `json:"intro" bson:"intro"`
	Content    string        `json:"content" bson:"content"`
	CoverImage string        `json:"cover_image" bson:"cover_image"`
	Category   Category      `json:"category" bson:"category"`
	CreatedOn  time.Time     `json:"created_on" bson:"created_on"`
	UpdatedOn  time.Time     `json:"updated_on" bson:"updated_on"`
}

//Activities is a slices of Activity
type Activities []Activity
