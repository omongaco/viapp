package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

//Category is a parent of Activity
type Category struct {
	ID         bson.ObjectId `json:"id" bson:"_id"`
	Name       string        `json:"name" bson:"name"`
	Content    string        `json:"content" bson:"content"`
	CoverImage string        `json:"cover_image" bson:"cover_image"`
	CreatedOn  time.Time     `json:"created_on" bson:"created_on"`
	UpdatedOn  time.Time     `json:"updated_on" bson:"updated_on"`
}

//Categories is a slices of Category
type Categories []Category
