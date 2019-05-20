

package authorization

import (
	"log"
	"net"
	"os"

	"github.com/user/2019_1_newTeam2/pkg/config"
	"github.com/user/2019_1_newTeam2/pkg/logger"
	"google.golang.org/grpc"
)

type AuthServer struct {
	ServerConfig *config.Config
	Logger       logger.LoggerInterface
	rpcServer    *grpc.Server
	AuthClient   AuthCheckerClient
}

func NewServer(pathToConfig string) (*AuthServer, error) {
	server := new(AuthServer)

	logger := new(logger.GoLogger)
	logger.SetOutput(os.Stderr)
	logger.SetPrefix("AUTH LOG: ")
	server.Logger = logger

	newConfig, err := config.NewConfig(pathToConfig)
	if err != nil {
		server.Logger.Log(err)
		return nil, err
	}
	server.ServerConfig = newConfig

	server.rpcServer = grpc.NewServer()
	RegisterAuthCheckerServer(server.rpcServer, NewAuthManager())
	return server, nil
}

func (server *AuthServer) Run() {
	lis, err := net.Listen("tcp", ":"+server.ServerConfig.Port)
	if err != nil {
		server.Logger.Logf("Can`t listen port %s", server.ServerConfig.Port)
		return
	}
	server.Logger.Logf("Running AuthMS(grps) on port %s", server.ServerConfig.Port)
	log.Fatal(server.rpcServer.Serve(lis))
}
