package controllers

import (
	"api-rest/database"
	"api-rest/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home")
}

func AllPersons(w http.ResponseWriter, r *http.Request) {
	var persons []models.Person
	database.DB.Table("persons").Find(&persons)
	json.NewEncoder(w).Encode(persons)
}

func GetPerson(w http.ResponseWriter, r *http.Request) {
	var persons models.Person
	vars := mux.Vars(r)
	id := vars["id"]
	database.DB.Table("persons").Where("id = ?", id).Find(&persons)
	json.NewEncoder(w).Encode(persons)
}

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	var person models.Person
	json.NewDecoder(r.Body).Decode(&person)
	database.DB.Table("persons").Create(&person)
	json.NewEncoder(w).Encode(person)
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
	var person models.Person
	vars := mux.Vars(r)
	id := vars["id"]

	database.DB.Table("persons").Delete(&person, id)
	json.NewEncoder(w).Encode(person)
}

func EditPerson(w http.ResponseWriter, r *http.Request) {
	var person models.Person
	vars := mux.Vars(r)
	id := vars["id"]

	database.DB.Table("persons").First(&person, id)
	json.NewDecoder(r.Body).Decode(&person)
	database.DB.Table("persons").Save(&person)

	json.NewEncoder(w).Encode(person)
}
