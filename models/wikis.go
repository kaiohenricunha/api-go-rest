package models

type Wiki struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Bio  string `json:"bio"`
}

var Wikis []Wiki
