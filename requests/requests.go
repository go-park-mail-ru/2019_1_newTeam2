package requests

type UserTableElem struct {
	Username string `json:"username"`
	Score    int    `json:"score, int"`
}

type UserAuth struct {
	Username string `json:"login"`
	Password string `json:"password"`
}

type Error struct {
	ThisIsERROR string `json:"error"`
}
