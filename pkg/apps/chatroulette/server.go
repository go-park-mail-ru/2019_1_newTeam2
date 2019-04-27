package chatroulette

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

func RunMePlease() {
	lis, err := net.Listen("tcp", ":8091")
	if err != nil {
		log.Fatalln("cant listet port", err)
	}

	server := grpc.NewServer()

	RegisterChatrouletteServer(server, NewChatManager())

	fmt.Println("starting server at :8091")
	server.Serve(lis)
}
