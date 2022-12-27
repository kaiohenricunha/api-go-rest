package main

import (
	"fmt"

	"github.com/kaiohenricunha/api-go-rest/models"
	"github.com/kaiohenricunha/api-go-rest/routes"
)

func main() {
	models.Wikis = []models.Wiki{
		{Id: 1, Name: "Nome 1", Bio: "Historia 1"},
		{Id: 2, Name: "Nome 2", Bio: "Historia 2"},
	}

	fmt.Println("Starting rest api server...")
	routes.HandleResquest()
}
