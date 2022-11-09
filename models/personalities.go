package models

type Personality struct {
	ID   int    `json:"id"`
	Name     string `json:"name"`
	Bio string `json:"bio"`
}

var Personalities []Personality

Personalities = []Personality{
	Personality{Name: "Kaio Henrique", Bio: "I'm a DevOps Engineer"},
	Personality{Name: "John Doe", Bio: "I'm a software developer"},
}
