package models

type Projects struct {
	Name string `json:"Name"`
	Description []string `json:"Description"`
	Technologies []string `json:"Technologies"`
}