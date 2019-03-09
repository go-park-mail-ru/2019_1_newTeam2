package requests

type UserTableElem struct {
	Username string `json:"username"`
	Score    int    `json:"score, int"`
}

type UserAuth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
