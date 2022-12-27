package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kaiohenricunha/api-go-rest/controllers"
)

func HandleResquest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/wikis", controllers.AllWikis).Methods("Get")
	r.HandleFunc("/api/wikis/{id}", controllers.ReturnAWiki).Methods("Get")
	log.Fatal(http.ListenAndServe(":8000", r))
}
