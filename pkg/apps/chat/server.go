package chat

import (
	"github.com/user/2019_1_newTeam2/pkg/middlewares"
	"github.com/user/2019_1_newTeam2/pkg/apps/chat/wshub"
	"log"
	"net/http"
	"os"

	"google.golang.org/grpc"

	"github.com/gorilla/mux"
	"github.com/user/2019_1_newTeam2/pkg/config"
	"github.com/user/2019_1_newTeam2/pkg/logger"
	"github.com/user/2019_1_newTeam2/storage"
	"github.com/user/2019_1_newTeam2/storage/interfaces"
	"github.com/user/2019_1_newTeam2/pkg/apps/authorization"
)

type ChatServer struct {
	Router       *mux.Router
	ServerConfig *config.Config
	Logger       logger.LoggerInterface
	Hub          wshub.IWSCommunicator
	CookieField  string
	AuthClient	 authorization.AuthCheckerClient
	DB         interfaces.DBChatInterface
}

func NewChatServer(pathToConfig string) (*ChatServer, error) {
	server := new(ChatServer)
	logger := new(logger.GoLogger)
	logger.SetOutput(os.Stderr)
	logger.SetPrefix("WORLDCHAT LOG: ")
	server.Logger = logger
	newConfig, err := config.NewConfig(pathToConfig)
	if err != nil {
		return nil, err
	}
	server.ServerConfig = newConfig
	newDB, err := storage.NewDataBase(server.ServerConfig.DBUser, server.ServerConfig.DBPassUser)
	if err != nil {
		return nil, err
	}
	server.DB = newDB
	server.CookieField = "session_id"
	server.Hub = wshub.NewWSCommunicator(server.ServerConfig.DBUser, server.ServerConfig.DBPassUser)
	router := mux.NewRouter()
	router.Use(middlewares.CreateCorsMiddleware(server.ServerConfig.AllowedHosts))
	router.Use(middlewares.CreateLoggingMiddleware(os.Stdout, "Word Trainer"))
	router.Use(middlewares.CreatePanicRecoveryMiddleware())

	chatRouter := router.PathPrefix("/chat/").Subrouter()
	chatRouter.HandleFunc("/enter/{id:[0-9]+}", server.CreateChat)
	chatRouter.HandleFunc("/history/", server.GetHistory)
	server.Router = router
	return server, nil
}

func (server *ChatServer) Run() {
	grcpAuthConn, err := grpc.Dial(
		"127.0.0.1:8092",
		grpc.WithInsecure(),
	)
	if err != nil {
		server.Logger.Log("Can`t connect ro grpc (auth ms)")
	}
	defer grcpAuthConn.Close()
	server.AuthClient = authorization.NewAuthCheckerClient(grcpAuthConn)

	port := server.ServerConfig.Port
	server.Logger.Logf("Running app on port %s", port)
	log.Fatal(http.ListenAndServe(":" + port, server.Router))
}