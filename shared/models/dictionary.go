package models

type CreateDictionary struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Cards       []Card `json:"cards,omitempty"`
}

//easyjson:json
type DictInfos []DictionaryInfo

//easyjson:json
type DictionaryInfo struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name"`
	Description string `json:"description"`
	UserId      int    `json:"userId,omitempty"`
}

type DictionaryInfoPrivilege struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	UserId      int    `json:"userId"`
	Privilege   bool   `json:"privilege,omitempty"`
}

type ParametersId struct {
	ID int `json:"id"`
}
