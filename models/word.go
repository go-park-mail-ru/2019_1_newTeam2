package models

type Word1 struct {
	Name     string    `json:"name,omitempty"`
	Language *Language `json:"language,omitempty"`
}

type AddedToDictWord struct {
	Name     string `json:"name,omitempty"`
	Language int    `json:"langID,omitempty"`
}
