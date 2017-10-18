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
type CategoryController struct {
	session *mgo.Session
}

//NewDestinationController
func NewCategoryController(s *mgo.Session) *CategoryController {
	return &CategoryController{s}
}

//GetDestinations finds all destinations in DB
func (dc CategoryController) GetCategories(w http.ResponseWriter, r *http.Request) {
	c := dc.session.DB("visitindonesiadb").C("category")
	category := models.Categories{}

	if err := c.Find(nil).All(&category); err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	uj, _ := json.Marshal(category)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", uj)
}

//GetDestination return a single Destination
func (dc CategoryController) GetCategory(w http.ResponseWriter, r *http.Request) {
	c := dc.session.DB("visitindonesiadb").C("category")
	vars := mux.Vars(r)
	id := vars["id"]
	category := models.Category{}

	if err := c.FindId(bson.ObjectIdHex(id)).One(&category); err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	uj, _ := json.Marshal(category)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj)
}

func (dc CategoryController) CreateCategory(w http.ResponseWriter, r *http.Request) {
	c := dc.session.DB("visitindonesiadb").C("category")
	category := models.Category{}

	err := json.NewDecoder(r.Body).Decode(&category)
	category.ID = bson.NewObjectId()
	category.CreatedOn = time.Now()
	category.UpdatedOn = time.Now()

	if err = c.Insert(category); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	uj, _ := json.Marshal(category)

	w.Header().Set("Centent-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s", uj)
}

func (dc CategoryController) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	c := dc.session.DB("visitindonesiadb").C("category")
	vars := mux.Vars(r)
	id := vars["id"]
	category := models.Category{}

	err := json.NewDecoder(r.Body).Decode(&category)

	if err = c.UpdateId(bson.ObjectIdHex(id), &category); err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	uj, _ := json.Marshal(category)

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", uj)
}

func (dc CategoryController) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	c := dc.session.DB("visitindonesiadb").C("category")
	vars := mux.Vars(r)
	id := vars["id"]

	if err := c.RemoveId(bson.ObjectIdHex(id)); err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
