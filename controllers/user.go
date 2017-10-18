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

//ProvinceController is a struct that points to mgo.Session
type ProvinceController struct {
	session *mgo.Session
}

//NewProvinceController to create new province object
func NewProvinceController(s *mgo.Session) *ProvinceController {
	return &ProvinceController{s}
}

//GetProvinces finds all provinces in DB
func (dc ProvinceController) GetProvinces(w http.ResponseWriter, r *http.Request) {
	c := dc.session.DB("visitindonesiadb").C("province")
	province := models.Provinces{}

	if err := c.Find(nil).All(&province); err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	uj, _ := json.Marshal(province)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", uj)
}

//GetProvince return a single Province
func (dc ProvinceController) GetProvince(w http.ResponseWriter, r *http.Request) {
	c := dc.session.DB("visitindonesiadb").C("province")
	vars := mux.Vars(r)
	id := vars["id"]
	province := models.Province{}

	if err := c.FindId(bson.ObjectIdHex(id)).One(&province); err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	uj, _ := json.Marshal(province)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj)
}

func (dc ProvinceController) CreateProvince(w http.ResponseWriter, r *http.Request) {
	c := dc.session.DB("visitindonesiadb").C("province")
	province := models.Province{}

	err := json.NewDecoder(r.Body).Decode(&province)
	province.ID = bson.NewObjectId()
	province.CreatedOn = time.Now()
	province.UpdatedOn = time.Now()

	if err = c.Insert(province); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	uj, _ := json.Marshal(province)

	w.Header().Set("Centent-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s", uj)
}

func (dc ProvinceController) UpdateProvince(w http.ResponseWriter, r *http.Request) {
	c := dc.session.DB("visitindonesiadb").C("province")
	vars := mux.Vars(r)
	id := vars["id"]
	province := models.Province{}

	err := json.NewDecoder(r.Body).Decode(&province)

	if err = c.UpdateId(bson.ObjectIdHex(id), &province); err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	uj, _ := json.Marshal(province)

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", uj)
}

func (dc ProvinceController) DeleteProvince(w http.ResponseWriter, r *http.Request) {
	c := dc.session.DB("visitindonesiadb").C("province")
	vars := mux.Vars(r)
	id := vars["id"]

	if err := c.RemoveId(bson.ObjectIdHex(id)); err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
