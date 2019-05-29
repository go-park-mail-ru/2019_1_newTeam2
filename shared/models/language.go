package models

//easyjson:json
type Langs []Language

//easyjson:json
type Language struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"username"`
}
