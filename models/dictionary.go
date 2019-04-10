package models

type Dictionary struct {
	ID       int       `json:"id,omitempty"`
	Name     string    `json:"name,omitempty"`
	Language *Language `json:"language,omitempty"`
	Cards    []Card    `json:"cards,omitempty"`
}
