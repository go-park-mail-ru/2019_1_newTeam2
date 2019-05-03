package main

import (
	"github.com/user/2019_1_newTeam2/pkg/apps/game"
	"os"
)

func main() {
	pathToConfig := ""
	if len(os.Args) != 2 {
		panic("Usage: ./main <path_to_config>")
	} else {
		pathToConfig = os.Args[1]
	}

	server, err := game.NewGameServer(pathToConfig)
	if err != nil {
		panic(err)
	}
	server.Run()
}
