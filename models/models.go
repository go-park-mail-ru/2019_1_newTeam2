package models

type User struct {
	ID          int    `json:"id,omitempty"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	Password    string `json:"password,omitempty"`
	LangID      int    `json:"langID, int"`
	PronounceON int    `json:"pronounceOn, int"`
	Score       int    `json:"score, int"`
	AvatarPath  string `json:"path"`
}

type UserTableElem struct {
	Username string `json:"username"`
	Score    int    `json:"score, int"`
}

type UserAuth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Error struct {
	ThisIsERROR string `json:"error"`
}
