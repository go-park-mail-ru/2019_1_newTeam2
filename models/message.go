package models

type GameMessage struct {
	Type string `json:"type"`
	Payload interface{} `json:"payload,omitempty"`
}