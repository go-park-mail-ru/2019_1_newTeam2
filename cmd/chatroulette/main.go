package main

import (
	"github.com/user/2019_1_newTeam2/pkg/apps/chatroulette"
)


func main() {
	server, err := chatroulette.NewChatServer()
	if err != nil {
		panic(err)
	}
	server.Run()
}