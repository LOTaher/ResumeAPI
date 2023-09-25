package models


type Skills struct {
	ID int `json:"ID,omitempty"`
	Languages []string `json:"Languages"`
	Frameworks []string `json:"Frameworks"`
	DeveloperTools []string `json:"Tools"`
	Libraries []string `json:"Libraries"`
}