package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/omongaco/viapp/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//DestinationController is a struct that points to mgo.Session
type ActivityController struct {
	session *mgo.Session
}

//NewDestinationController
func NewActivityController(s *mgo.Session) *ActivityController {
	return &ActivityController{s}
}

//GetDestinations finds all destinations in DB
func (dc ActivityController) GetActivities(w http.ResponseWriter, r *http.Request) {
	c := dc.session.DB("visitindonesiadb").C("activity")
	activity := models.Activities{}

	if err := c.Find(nil).All(&activity); err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	uj, _ := json.Marshal(activity)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", uj)
}

//GetDestination return a single Destination
func (dc ActivityController) GetActivity(w http.ResponseWriter, r *http.Request) {
	c := dc.session.DB("visitindonesiadb").C("activity")
	vars := mux.Vars(r)
	id := vars["id"]
	activity := models.Activity{}

	if err := c.FindId(bson.ObjectIdHex(id)).One(&activity); err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	uj, _ := json.Marshal(activity)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj)
}

func (dc ActivityController) CreateActivity(w http.ResponseWriter, r *http.Request) {
	c := dc.session.DB("visitindonesiadb").C("activity")
	activity := models.Activity{}

	err := json.NewDecoder(r.Body).Decode(&activity)
	activity.ID = bson.NewObjectId()
	activity.CreatedOn = time.Now()
	activity.UpdatedOn = time.Now()

	if err = c.Insert(activity); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	uj, _ := json.Marshal(activity)

	w.Header().Set("Centent-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s", uj)
}

func (dc ActivityController) UpdateActivity(w http.ResponseWriter, r *http.Request) {
	c := dc.session.DB("visitindonesiadb").C("activity")
	vars := mux.Vars(r)
	id := vars["id"]
	activity := models.Activity{}

	err := json.NewDecoder(r.Body).Decode(&activity)

	if err = c.UpdateId(bson.ObjectIdHex(id), &activity); err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	uj, _ := json.Marshal(activity)

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", uj)
}

func (dc ActivityController) DeleteActivity(w http.ResponseWriter, r *http.Request) {
	c := dc.session.DB("visitindonesiadb").C("activity")
	vars := mux.Vars(r)
	id := vars["id"]

	if err := c.RemoveId(bson.ObjectIdHex(id)); err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
