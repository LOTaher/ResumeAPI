package models

type Projects struct {
	ID int `json:"ID,omitempty"`
	Name string `json:"Name"`
	Description []string `json:"Description"`
	Technologies []string `json:"Technologies"`
}