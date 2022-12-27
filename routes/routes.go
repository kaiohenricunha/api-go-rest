package routes

import (
	"log"
	"net/http"

	"github.com/kaiohenricunha/go-rest-api/controllers"
)

func HandleResquest() {
	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/api/wikis", controllers.AllWikis)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
