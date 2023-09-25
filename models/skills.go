package models


type Skills struct {
	ID int `json:"ID"`
	Languages []string `json:"Languages"`
	Frameworks []string `json:"Frameworks"`
	DeveloperTools []string `json:"Tools"`
	Libraries []string `json:"Libraries"`
}