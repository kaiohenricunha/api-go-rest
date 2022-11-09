package routes

import (
	"log"
	"net/http"

	"github.com/kaiohenricunha/api-go-rest/controllers"
)

func HandleResquest() {
	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/api/personality", controllers.AllPersonalities)
	log.Fatal(http.ListenAndServe(":8000", nil))
}