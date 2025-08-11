package models

type Currency struct {
	Code   string `json:"code"`
	Name   string `json:"name,omitempty"`
	Symbol string `json:"symbol,omitempty"`
	// Add other fields if your actual model has more
}
