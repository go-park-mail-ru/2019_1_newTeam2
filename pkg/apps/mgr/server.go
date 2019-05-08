package mgr

import (
	"log"
	"net"
	"os"

	"github.com/user/2019_1_newTeam2/pkg/config"
	"github.com/user/2019_1_newTeam2/pkg/logger"
	"google.golang.org/grpc"
)

type ScoreServer struct {
	ServerConfig *config.Config
	Logger       logger.LoggerInterface
	rpcServer    *grpc.Server
	UserScore    UserScoreUpdaterClient
}

func NewServer(pathToConfig string) (*ScoreServer, error) {
	server := new(ScoreServer)

	logger := new(logger.GoLogger)
	logger.SetOutput(os.Stderr)
	logger.SetPrefix("SCORE LOG: ")
	server.Logger = logger

	newConfig, err := config.NewConfig(pathToConfig)
	if err != nil {
		server.Logger.Log(err)
		return nil, err
	}
	server.ServerConfig = newConfig

	server.rpcServer = grpc.NewServer()
	RegisterUserScoreUpdaterServer(server.rpcServer, NewUserScoreUpdaterManager(server.ServerConfig.DBUser, server.ServerConfig.DBPassUser))
	return server, nil
}

func (server *ScoreServer) Run() {
	lis, err := net.Listen("tcp", ":"+server.ServerConfig.Port)
	if err != nil {
		server.Logger.Logf("Can`t listen port %s", server.ServerConfig.Port)
		return
	}
	server.Logger.Logf("Running ScoreMS(grps) on port %s", server.ServerConfig.Port)
	log.Fatal(server.rpcServer.Serve(lis))
}
