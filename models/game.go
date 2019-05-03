package models

type GameResult struct {
	ID      int  `json:"id"`
	Correct bool `json:"correct"`
}

type GameResults []GameResult

type GameWord struct {
	CardId   int      `json:"id"`
	Word     string   `json:"word"`
	Variants []string `json:"variants,omitempty"`
	Correct  int      `json:"correct"`
}
