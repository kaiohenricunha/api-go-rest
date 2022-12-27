package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/kaiohenricunha/api-go-rest/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func AllWikis(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Wikis)
}

func ReturnAWiki(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for _, wiki := range models.Wikis {
		if strconv.Itoa(wiki.Id) == id {
			json.NewEncoder(w).Encode(wiki)
		}
	}
}
