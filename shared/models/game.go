package models

//easyjson:json
type GameResult struct {
	ID      int  `json:"id"`
	Correct bool `json:"correct"`
}

//easyjson:json
type GameResults []GameResult

//easyjson:json
type GameWords []GameWord

//easyjson:json
type GameWord struct {
	CardId   int      `json:"id"`
	Word     string   `json:"word"`
	Variants []string `json:"variants,omitempty"`
	Correct  int      `json:"correct"`
}

//easyjson:json
type GameQuestions []GameQuestion


//easyjson:json
type GameQuestion struct {
	Words    [4]string `json:"words"`
	Answer   string    `json:"answer"`
	Question string    `json:"question"`
}
