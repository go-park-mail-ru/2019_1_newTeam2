package models

type CreateDictionary struct {
	ID          int               `json:"id,omitempty"`
	Name        string            `json:"name,omitempty"`
	Description string            `json:"description,omitempty"`
	Cards       []AddedToDictCard `json:"cards,omitempty"`
}

type DictionaryInfo struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}
