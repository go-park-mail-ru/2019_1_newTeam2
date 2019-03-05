package main

type userAuthorization struct {
	username string `json:"username"`
	password string `json:"password"`
}

type userRegistration struct {
	username string `json:"username"`
	email    string `json:"email"`
	password string `json:"password"`
}

type userUpdate struct {
	username string `json:"username,omitempty"`
	email    string `json:"email,omitempty"`
	password string `json:"password,omitempty"`
}
