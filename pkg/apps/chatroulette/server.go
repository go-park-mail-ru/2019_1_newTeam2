package chatroulette

import (
	"github.com/user/2019_1_newTeam2/pkg/middlewares"
	"github.com/user/2019_1_newTeam2/pkg/wshub"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"github.com/user/2019_1_newTeam2/pkg/config"
	"github.com/user/2019_1_newTeam2/pkg/logger"
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

	newConfig, err := config.NewConfig(pathToConfig)
	if err != nil {
		return nil, err
	}

	server.CookieField = "session_id"

	server.ServerConfig = newConfig
	server.Hub = wshub.NewWSCommunicator()
	router := mux.NewRouter()
	router.Use(middlewares.CreateCorsMiddleware(server.ServerConfig.AllowedHosts))
	router.Use(middlewares.CreateLoggingMiddleware(os.Stdout, "Word Trainer"))
	router.Use(middlewares.CreatePanicRecoveryMiddleware())
	chatRouter := router.PathPrefix("/chat/").Subrouter()
	chatRouter.HandleFunc("/enter/", server.CreateChat)
	chatRouter.HandleFunc("/history/", server.GetHistory)
	server.Router = router
	return server, nil
}

func (server *ChatServer) Run() {
	port := os.Getenv("PORT")		// change for necessary port
	if port == "" {
		port = "8091"
	}
	log.Println("running")
	log.Fatal(http.ListenAndServe(":" + port, server.Router))
}