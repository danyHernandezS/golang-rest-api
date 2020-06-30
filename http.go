package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//type person
type Person struct {
	Nombre       string `json:"nombre,omitempty"`
	Departamento string `json:"departamento,omitempty"`
	Edad         int    `json:"edad,omitempty"`
	Estado       string `json:"estado,omitempty"`
	Contagio     string `json:"forma de contagio,omitempty"`
}

var people []Person

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Wecome the my GO API!")
}

// get method
func GetPeopleEndpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(people)
}

//post method
func CreatePersonEndpoint(w http.ResponseWriter, req *http.Request) {
	var person Person
	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Fprintf(w, "Insert a Valid Task Data")
	}
	json.Unmarshal(reqBody, &person)
	people = append(people, person)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(person)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	// endpoints
	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")
	router.HandleFunc("/people", CreatePersonEndpoint).Methods("POST")

	log.Fatal(http.ListenAndServe(":3000", router))
}
