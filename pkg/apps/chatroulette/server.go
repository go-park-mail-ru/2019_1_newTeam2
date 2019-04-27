package chatroulette

import (
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
}

func NewChatServer() (*ChatServer, error) {
	server := new(ChatServer)
	router := mux.NewRouter()

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
	log.Fatal(http.ListenAndServe(":" + port, server.Router))
}