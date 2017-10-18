package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/omongaco/viapp/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//DestinationController is a struct that points to mgo.Session
type DestinationController struct {
	session *mgo.Session
}

//NewDestinationController
func NewDestinationController(s *mgo.Session) *DestinationController {
	return &DestinationController{s}
}

//GetDestinations finds all destinations in DB
func (dc DestinationController) GetDestinations(w http.ResponseWriter, r *http.Request) {
	c := dc.session.DB("visitindonesiadb").C("destination")
	destination := models.Destinations{}

	if err := c.Find(nil).All(&destination); err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	uj, _ := json.Marshal(destination)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", uj)
}

//GetDestination return a single Destination
func (dc DestinationController) GetDestination(w http.ResponseWriter, r *http.Request) {
	c := dc.session.DB("visitindonesiadb").C("destination")
	vars := mux.Vars(r)
	id := vars["id"]
	destination := models.Destination{}

	if err := c.FindId(bson.ObjectIdHex(id)).One(&destination); err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	uj, _ := json.Marshal(destination)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj)
}

func (dc DestinationController) CreateDestination(w http.ResponseWriter, r *http.Request) {
	c := dc.session.DB("visitindonesiadb").C("destination")
	destination := models.Destination{}

	err := json.NewDecoder(r.Body).Decode(&destination)
	destination.ID = bson.NewObjectId()

	if err = c.Insert(destination); err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	uj, _ := json.Marshal(destination)

	w.Header().Set("Centent-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s", uj)
}

func (dc DestinationController) UpdateDestination(w http.ResponseWriter, r *http.Request) {
	c := dc.session.DB("visitindonesiadb").C("destination")
	vars := mux.Vars(r)
	id := vars["id"]
	destination := models.Destination{}

	err := json.NewDecoder(r.Body).Decode(&destination)

	if err = c.UpdateId(bson.ObjectIdHex(id), &destination); err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	uj, _ := json.Marshal(destination)

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", uj)
}

func (dc DestinationController) DeleteDestination(w http.ResponseWriter, r *http.Request) {
	c := dc.session.DB("visitindonesiadb").C("destination")
	vars := mux.Vars(r)
	id := vars["id"]

	if err := c.RemoveId(bson.ObjectIdHex(id)); err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
