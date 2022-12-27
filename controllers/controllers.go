package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kaiohenricunha/go-rest-api/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func AllWikis(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Wikis)
}
