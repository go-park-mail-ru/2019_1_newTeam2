package models

type CreateDictionary struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Cards       []Card `json:"cards,omitempty"`
}

type DictionaryInfo struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	UserId 		int	   `json:"userId"`
}
