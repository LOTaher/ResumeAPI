package models

type Education struct {
	ID    int    `json:"ID,omitempty"`
	Year   string `json:"Year"`
	Degree string `json:"Degree"`
	School string `json:"School"`
}