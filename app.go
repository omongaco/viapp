package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/omongaco/viapp/controllers"
	mgo "gopkg.in/mgo.v2"
)

func main() {
	router := mux.NewRouter()

	aCon := controllers.NewActivityController(getSession())
	cCon := controllers.NewCategoryController(getSession())
	dCon := controllers.NewDestinationController(getSession())
	pCon := controllers.NewProvinceController(getSession())

	//Activity API Endpoint
	router.HandleFunc("/api/activity/", aCon.GetActivities).Methods("GET")
	router.HandleFunc("/api/activity/{id}", aCon.GetActivity).Methods("GET")
	router.HandleFunc("/api/activity/", aCon.CreateActivity).Methods("POST")
	router.HandleFunc("/api/activity/{id}", aCon.UpdateActivity).Methods("PUT")
	router.HandleFunc("/api/activity/{id}", aCon.DeleteActivity).Methods("DELETE")

	//Category API Endpoint
	router.HandleFunc("/api/category/", cCon.GetCategories).Methods("GET")
	router.HandleFunc("/api/category/{id}", cCon.GetCategory).Methods("GET")
	router.HandleFunc("/api/category/", cCon.CreateCategory).Methods("POST")
	router.HandleFunc("/api/category/{id}", cCon.UpdateCategory).Methods("PUT")
	router.HandleFunc("/api/category/{id}", cCon.DeleteCategory).Methods("DELETE")

	//Destination API Endpoint
	router.HandleFunc("/api/destination/", dCon.GetDestinations).Methods("GET")
	router.HandleFunc("/api/destination/{id}", dCon.GetDestination).Methods("GET")
	router.HandleFunc("/api/destination/", dCon.CreateDestination).Methods("POST")
	router.HandleFunc("/api/destination/{id}", dCon.UpdateDestination).Methods("PUT")
	router.HandleFunc("/api/destination/{id}", dCon.DeleteDestination).Methods("DELETE")

	//Province API Endpoint
	router.HandleFunc("/api/province/", pCon.GetProvinces).Methods("GET")
	router.HandleFunc("/api/province/{id}", pCon.GetProvince).Methods("GET")
	router.HandleFunc("/api/province/", pCon.CreateProvince).Methods("POST")
	router.HandleFunc("/api/province/{id}", pCon.UpdateProvince).Methods("PUT")
	router.HandleFunc("/api/province/{id}", pCon.DeleteProvince).Methods("DELETE")

	if err := http.ListenAndServe(":3000", router); err != nil {
		log.Fatal(err)
	}
}

func getSession() *mgo.Session {
	//Connect to MOngoDB
	s, err := mgo.Dial("localhost")

	if err != nil {
		panic(err)
	}

	return s
}
