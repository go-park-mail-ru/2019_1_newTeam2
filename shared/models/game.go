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

type GameQuestion struct {
	Words    [4]string `json:"words"`
	Answer   string    `json:"answer"`
	Question string    `json:"question"`
}
