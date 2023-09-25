package models

type Info struct {
	ID int `json:"ID,omitempty"`
	Name string `json:"Name"`
	Email string `json:"Email"`
	Phone string `json:"Phone"`
	Website string `json:"Website"`
	Availability string `json:"Availability"`
	Linkedin string `json:"Linkedin"`
	Github string `json:"Github"`
}