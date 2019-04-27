package models

type Message struct {
	ID          int       `json:"id,omitempty"`
	Data        string	  `json:"message"`
}

type Dialog struct {
	ID          int			 `json:"id,omitempty"`
	User1		int			 `json:"user1,omitempty"`
	User2		int			 `json:"user2,omitempty"`
	Messages    []Message    `json:"cards,omitempty"`
}