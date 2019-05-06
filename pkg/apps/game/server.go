package game

import (
	"net/http"
	"os"

	"google.golang.org/grpc"

	"github.com/gorilla/mux"
	"github.com/user/2019_1_newTeam2/pkg/apps/authorization"
	"github.com/user/2019_1_newTeam2/pkg/apps/game/game"
	"github.com/user/2019_1_newTeam2/pkg/config"
	"github.com/user/2019_1_newTeam2/pkg/logger"
	"github.com/user/2019_1_newTeam2/pkg/middlewares"
)

type GameServer struct {
	Router       *mux.Router
	ServerConfig *config.Config
	Logger       logger.LoggerInterface
	AuthClient   authorization.AuthCheckerClient
	CookieField  string
	Game         *game.Game
}

func NewGameServer(pathToConfig string) (*GameServer, error) {
	server := new(GameServer)
	logger := new(logger.GoLogger)
	logger.SetOutput(os.Stderr)
	logger.SetPrefix("GAME LOG: ")
	server.Logger = logger

	newConfig, err := config.NewConfig(pathToConfig)
	if err != nil {
		return nil, err
	}
	server.ServerConfig = newConfig

	server.CookieField = "session_id"

	router := mux.NewRouter()
	router = router.PathPrefix("/multiplayer/").Subrouter()
	router.Use(middlewares.CreateCorsMiddleware(server.ServerConfig.AllowedHosts))
	router.Use(middlewares.CreateLoggingMiddleware(os.Stdout, "Word Trainer"))
	router.Use(middlewares.CreatePanicRecoveryMiddleware())
	router.Use(middlewares.CreateCheckAuthMiddleware([]byte(server.ServerConfig.Secret), server.CookieField, server.IsLogined))

	router.HandleFunc("/game", server.OpenConnection)

	server.Router = router
	server.Game = game.NewGame(server.ServerConfig.DBUser, server.ServerConfig.DBPassUser)
	return server, nil
}

func (server *GameServer) Run() {
	grcpAuthConn, err := grpc.Dial(
		server.ServerConfig.AuthHost+":"+server.ServerConfig.AuthPort,
		grpc.WithInsecure(),
	)
	if err != nil {
		server.Logger.Log("Can`t connect ro grpc (auth ms)")
	}
	defer grcpAuthConn.Close()
	go server.Game.Run()
	server.AuthClient = authorization.NewAuthCheckerClient(grcpAuthConn)
	server.Logger.Logf("Running app on port %s", server.ServerConfig.Port)
	server.Logger.Log(http.ListenAndServe(":"+server.ServerConfig.Port, server.Router))
}
