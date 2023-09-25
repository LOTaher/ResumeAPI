package models

type Experience struct {
	ID   int    `json:"ID,omitempty"`
	Company string `json:"Company"`
	Position string `json:"Position"`
	StartDate string `json:"StartDate"`
	EndDate string `json:"EndDate"`
	Responsibilities []string `json:"Responsibilities"`
}