package main

import (
	"github.com/user/2019_1_newTeam2/server"
	"os"
)

func main() {
	pathToConfig := ""
	if len(os.Args) != 2 {
		panic("Usage: ./main <path_to_config>")
	} else {
		pathToConfig = os.Args[1]
	}

	serv, err := server.NewServer(pathToConfig)
	if err != nil {
		panic(err.Error())
	}
	serv.Run()
}
