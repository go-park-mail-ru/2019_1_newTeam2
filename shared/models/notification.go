package models

type DictionaryNote struct {
	BorrowerId     int    `json:"borrowerId,omitempty, int"`
	DictionaryName string `json:"dictionaryName,omitempty"`
}
