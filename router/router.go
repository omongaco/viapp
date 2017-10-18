package router

import (
	"log"

	. "github.com/omongaco/viapp/controllers"
	mgo "gopkg.in/mgo.v2"
)

var db *mgo.Database

type RouterAccess struct {
	Server   string
	Database string
}

func (r *RouterAccess) Connect() {
	session, err := mgo.Dial(r.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(r.Database)
}
