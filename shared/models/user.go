package models

type User struct {
	ID          int    `json:"id,omitempty"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	Password    string `json:"password,omitempty"`
	LangID      int    `json:"langID,int"`
	PronounceON int    `json:"pronounceOn,int"`
	Score       int    `json:"score,int"`
	AvatarPath  string `json:"path"`
}


//easyjson:json
type TableUsers []UserTableElem

//easyjson:json
type UserTableElem struct {
	Username string `json:"username"`
	Score    int    `json:"score,int"`
}
