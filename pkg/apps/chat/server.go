package chat

import (
	"github.com/user/2019_1_newTeam2/pkg/middlewares"
	"github.com/user/2019_1_newTeam2/pkg/apps/chat/wshub"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"github.com/user/2019_1_newTeam2/pkg/config"
	"github.com/user/2019_1_newTeam2/pkg/logger"
	"github.com/user/2019_1_newTeam2/storage"
	"github.com/user/2019_1_newTeam2/storage/interfaces"
)

type ChatServer struct {
	Router       *mux.Router
	DB           interfaces.DBChatInterface
	ServerConfig *config.Config
	Logger       logger.LoggerInterface
	Hub          wshub.IWSCommunicator
	CookieField  string
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
	server.CookieField = "session_id"
	server.ServerConfig = newConfig
	newDB, err := storage.NewDataBase(server.ServerConfig.DBUser, server.ServerConfig.DBPassUser)
	if err != nil {
		return nil, err
	}
	server.DB = newDB
	server.Hub = wshub.NewWSCommunicator()
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
	port := server.ServerConfig.Port
	server.Logger.Logf("Running app on port %s", port)
	log.Fatal(http.ListenAndServe(":" + port, server.Router))
}