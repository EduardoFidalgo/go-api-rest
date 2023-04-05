package routes

import (
	"api-rest/controllers"
	"api-rest/middleware"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	router := mux.NewRouter()
	router.Use(middleware.ContentTypeMiddleware)
	router.HandleFunc("/", controllers.Home)
	router.HandleFunc("/api/persons", controllers.AllPersons).Methods("Get")
	router.HandleFunc("/api/persons/{id}", controllers.GetPerson).Methods("Get")
	router.HandleFunc("/api/persons", controllers.CreatePerson).Methods("Post")
	router.HandleFunc("/api/persons/{id}", controllers.DeletePerson).Methods("Delete")
	router.HandleFunc("/api/persons/{id}", controllers.EditPerson).Methods("Put")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(router)))
}
