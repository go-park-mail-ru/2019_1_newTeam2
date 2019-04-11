package models


type WordFull struct {
	Name     string    `json:"name,omitempty"`
	Language *Language `json:"language,omitempty"`
}


type Word struct {
	Name       string `json:"name,omitempty"`
	LanguageId int    `json:"languageId,omitempty,int"`
}
