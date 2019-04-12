package models

type CardFull struct {
	ID          int       `json:"id,omitempty"`
	Word        *WordFull `json:"word,omitempty"`
	Translation *WordFull `json:"translation,omitempty"`
}

type Card struct {
	ID          int   `json:"id,omitempty"`
	Word        *Word `json:"word,omitempty"`
	Translation *Word `json:"translation,omitempty"`
	Frequency float64 `json:"frequency,omitempty"`
}
