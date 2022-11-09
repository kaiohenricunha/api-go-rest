package routes

import (
	"log"
	"net/http"

	"github.com/kaiohenricunha/api-go-rest/controllers"
)

func HandleResquest() {
	r = mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personality", controllers.AllPersonalities)
	log.Fatal(http.ListenAndServe(":8000", r))
}
