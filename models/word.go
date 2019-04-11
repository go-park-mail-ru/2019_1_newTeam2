package models

type Word1 struct {
	Name     string    `json:"name,omitempty"`
	Language *Language `json:"language,omitempty"`
}

type WordFull struct {
	Name     string    `json:"name,omitempty"`
	Language *Language `json:"language,omitempty"`
}

type AddedToDictWord struct {
	Name     string `json:"name,omitempty"`
	Language int    `json:"langID,omitempty"`
}

type Word struct {
	Name       string `json:"name,omitempty"`
	LanguageId int    `json:"languageId,omitempty,int"`
}
