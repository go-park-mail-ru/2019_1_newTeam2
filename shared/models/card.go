package models

type CardFull struct {
	ID          int       `json:"id,omitempty"`
	Word        *WordFull `json:"word"`
	Translation *WordFull `json:"translation"`
}

//easyjson:json
type Cards []Card

//easyjson:json
type Card struct {
	ID          int     `json:"id,omitempty"`
	Word        *Word   `json:"word"`
	Translation *Word   `json:"translation"`
	Frequency   float64 `json:"frequency"`
}

type CardDelete struct {
	DictionaryId int `json:"dictionaryId"`
	CardId       int `json:"cardId"`
}
