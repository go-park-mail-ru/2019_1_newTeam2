package game

import (
	"net/http"
	"os"

	"google.golang.org/grpc"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/user/2019_1_newTeam2/pkg/apps/authorization"
	"github.com/user/2019_1_newTeam2/pkg/apps/game/game"
	"github.com/user/2019_1_newTeam2/pkg/apps/game/room"
	"github.com/user/2019_1_newTeam2/pkg/apps/mgr"
	"github.com/user/2019_1_newTeam2/pkg/config"
	"github.com/user/2019_1_newTeam2/pkg/logger"
	"github.com/user/2019_1_newTeam2/pkg/middlewares"
)

type GameServer struct {
	Router       *mux.Router
	ServerConfig *config.Config
	Logger       logger.LoggerInterface
	AuthClient   authorization.AuthCheckerClient
	ScoreClient  mgr.UserScoreUpdaterClient
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

	r := mux.NewRouter()
	router := r.PathPrefix("/multiplayer/").Subrouter()
	router.Use(middlewares.CreateCorsMiddleware(server.ServerConfig.AllowedHosts))
	router.Use(middlewares.CreateLoggingMiddleware(os.Stdout, "Word Trainer"))
	router.Use(middlewares.CreatePanicRecoveryMiddleware())
	router.Use(middlewares.CreateCheckAuthMiddleware([]byte(server.ServerConfig.Secret), server.CookieField, server.IsLogined))

	router.HandleFunc("/game", promhttp.InstrumentHandlerCounter(
		MultiplayerHitsMetric,
		http.HandlerFunc(server.OpenConnection),
	))

	r.Handle("/metrics", promhttp.Handler())
	server.Router = r
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

	grcpScoreConn, err := grpc.Dial(
		server.ServerConfig.ScoreHost+":"+server.ServerConfig.ScorePort,
		grpc.WithInsecure(),
	)
	if err != nil {
		server.Logger.Log("Can`t connect ro grpc (score ms)")
	}
	defer grcpScoreConn.Close()

	prometheus.MustRegister(MultiplayerHitsMetric, room.PlayerCountMetric, game.RoomCountMetric)

	server.AuthClient = authorization.NewAuthCheckerClient(grcpAuthConn)
	server.ScoreClient = mgr.NewUserScoreUpdaterClient(grcpScoreConn)
	server.Game = game.NewGame(server.ServerConfig.DBHost,
		server.ServerConfig.DBUser, server.ServerConfig.DBPassUser, server.ScoreClient)
	go server.Game.Run()
	server.Logger.Logf("Running app on port %s", server.ServerConfig.Port)
	server.Logger.Log(http.ListenAndServe(":"+server.ServerConfig.Port, server.Router))
}
