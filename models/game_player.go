package models

type PlayerData struct {
	Username string `json:"username"`
	Score	 int	`json:"score"`
}

type State struct {
	Players []PlayerData `json:"players"`
}