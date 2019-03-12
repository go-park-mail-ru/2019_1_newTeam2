package main

import (
	"os"
	"github.com/user/2019_1_newTeam2/server"
)

func main() {
	pathToConfig := ""
	if len(os.Args) != 2 {
		pathToConfig = "./config/config.json"
	} else {
		pathToConfig = os.Args[1]
	}

	serv, err := server.NewServer(pathToConfig)
	if err != nil {
		panic(err.Error())
	}
	serv.Run()
}