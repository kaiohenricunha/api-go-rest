package models

type Personality struct {
	Name     string `json:"name"`
	Bio string `json:"bio"`
}

var Personalities []Personality