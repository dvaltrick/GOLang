package main

import (
	"fmt"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/teste/models"
	"github.com/teste/services"
)

var uniList []models.University

const (
	dbuser = "postgres"
	dbpwd = "admin" 
	dbname = "goteste" 
)

func main() {
	//START APPLICATION 
	fmt.Println("Teste aptidão - Diego Ecker Valtrick")
	fmt.Println("Desenvolvimento com GO Language! ")

	//Verification to define if it's necessary to import data from a CSV file to a database
	uniList = services.GetAll()
	if len(uniList) == 0 { //if list is empty call a function to import database
		services.ImportFromFile()
	}

	//Create a route to HTTP access, set call function and define a method
	router := mux.NewRouter()
	router.HandleFunc("/universities", GetUniversities).Methods("POST")
	log.Fatal(http.ListenAndServe(":12345", router))

}

//Method of endpoint
func GetUniversities(w http.ResponseWriter, req *http.Request){
	var university models.University
	var uniList []models.University
	
	_ = json.NewDecoder(req.Body).Decode(&university)

	//Call a function to get data in the backend with the received parameters of frontend
	uniList = services.Get(university)

	//Define the limit of hundred objects to return
	uniList = uniList[0:100]
	json.NewEncoder(w).Encode(uniList)
}
