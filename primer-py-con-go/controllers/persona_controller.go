package controllers

import (
	"crud-api-rest-go/commons"
	"crud-api-rest-go/models"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func GetAll(writer http.ResponseWriter, request *http.Request) {
	personas := []models.Persona{}
	db := commons.GetConnection()
	// defer db.Close()

	db.Find(&personas)

	//con el _ ignoro la devolucion del error en la siguiente linea
	json, _ := json.Marshal(personas)

	commons.SendResponse(writer, http.StatusOK, json)
}

func Get(writer http.ResponseWriter, request *http.Request) {
	persona := models.Persona{}

	id := mux.Vars(request)["id"]

	db := commons.GetConnection()
	// defer db.Close()

	db.Find(&persona, id)

	if persona.ID > 0 {
		json, _ := json.Marshal(persona)
		commons.SendResponse(writer, http.StatusOK, json)
	} else {
		commons.SendError(writer, http.StatusNoContent)
	}
}

func Save(writer http.ResponseWriter, request *http.Request) {
	persona := models.Persona{}

	db := commons.GetConnection()
	// defer db.Close()

	error := json.NewDecoder(request.Body).Decode(&persona)

	if error != nil {
		log.Fatal(error)
		commons.SendError(writer, http.StatusBadRequest)
		return
	}

	error = db.Save(&persona).Error

	if error != nil {
		log.Fatal(error)
		commons.SendError(writer, http.StatusInternalServerError)
		return
	}

	json, _ := json.Marshal(persona)

	commons.SendResponse(writer, http.StatusCreated, json)
}

func Delete(writer http.ResponseWriter, request *http.Request) {
	persona := models.Persona{}

	id := mux.Vars(request)["id"]

	db := commons.GetConnection()
	// defer db.Close()

	db.Find(&persona, id)

	if persona.ID > 0 {
		db.Delete(persona)
		commons.SendResponse(writer, http.StatusOK, []byte(`{}`))
	} else {
		commons.SendError(writer, http.StatusNoContent)
	}
}
