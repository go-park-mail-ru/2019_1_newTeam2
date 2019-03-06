package requests

type UserTableElem struct {
	Username string `json:"username"`
	Score    int    `json:"score, int"`
}
