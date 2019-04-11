package models

type Card struct {
	ID          int   `json:"id,omitempty"`
	Word        *Word `json:"word,omitempty"`
	Translation *Word `json:"translation,omitempty"`
}

type AddedToDictCard struct {
	ID          int              `json:"id,omitempty"`
	Word        *AddedToDictWord `json:"word,omitempty"`
	Translation *AddedToDictWord `json:"translation,omitempty"`
}
