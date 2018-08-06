//resource https://www.thepolyglotdeveloper.com/2017/03/bundle-html-css-javascript-served-golang-application/
package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux" //make endpoints easier to create, more features
)

//omitempty: if the property is null, it will be excluded from the JSON data rather than showing up as an empty string or value.
type Person struct {
	ID        string   `json:"id,omitempty"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"` //Address need to be a pointer, otherwise omniempty will not work
}

type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var people []Person //slice of Person

func GetPeopleEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item) // found, display it
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{}) //not found, create an empty JSON object
}

func CreatePersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var person Person
	_ = json.NewDecoder(req.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people, person)
	json.NewEncoder(w).Encode(person)
}

func DeletePersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for i, item := range people {
		if item.ID == params["id"] {
			people = append(people[:i], people[i+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(people)
}

func main() {
	router := mux.NewRouter()
	people = append(people, Person{ID: "1", Firstname: "Nic", Lastname: "Raboy", Address: &Address{City: "Dublin", State: "CA"}})
	people = append(people, Person{ID: "2", Firstname: "Maria", Lastname: "Raboy"})
	router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", GetPeopleEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePersonEndpoint).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePersonEndpoint).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}
