package models

type Card struct {
	ID          int   `json:"id,omitempty"`
	Word        *Word `json:"word,omitempty"`
	Translation *Word `json:"translation,omitempty"`
}


