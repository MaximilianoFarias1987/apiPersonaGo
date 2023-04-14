package routes

import (
	"crud-api-rest-go/controllers"

	"github.com/gorilla/mux"
)

func SetPersonasRoutes(router *mux.Router) {
	subRoute := router.PathPrefix("/persona/api").Subrouter()

	subRoute.HandleFunc("/all", controllers.GetAll).Methods("GET")
	subRoute.HandleFunc("/save", controllers.Save).Methods("POST")
	subRoute.HandleFunc("/find/{id}", controllers.Get).Methods("GET")
	subRoute.HandleFunc("/delete/{id}", controllers.Delete).Methods("POST")
}
