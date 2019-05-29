package models


//easyjson:json
type GameMessages []GameMessage


//easyjson:json
type GameMessage struct {
	Type    string      `json:"type"`
	Payload interface{} `json:"payload,omitempty"`
}
