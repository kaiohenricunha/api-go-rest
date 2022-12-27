package models

type Wiki struct {
	Name     string `json:"name"`
	Bio string `json:"bio"`
}

var Wikis []Wiki