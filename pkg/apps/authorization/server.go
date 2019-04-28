package authorization

import (
	//"github.com/user/2019_1_newTeam2/pkg/apps/common"
	//"github.com/user/2019_1_newTeam2/pkg/wshub"
	"log"
	"net/http"
	"os"
	//"path/filepath"

	"github.com/gorilla/mux"
	grpc "github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
	"github.com/user/2019_1_newTeam2/pkg/config"
	"github.com/user/2019_1_newTeam2/pkg/logger"
	"github.com/user/2019_1_newTeam2/pkg/middlewares"
	"github.com/user/2019_1_newTeam2/storage"
	"github.com/user/2019_1_newTeam2/storage/interfaces"
)

type AuthServer struct {
	Router       *mux.Router
	DB           interfaces.DBInterface
	ServerConfig *config.Config
	Logger       logger.LoggerInterface
	CookieField  string
}

func NewServer(pathToConfig string) (*AuthServer, error) {
	server := new(AuthServer)

	logger := new(logger.GoLogger)
	logger.SetOutput(os.Stderr)
	logger.SetPrefix("AUTH LOG: ")
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

	router := mux.NewRouter()

	muxServer := grpc.NewServer()
	muxServer.RegisterCodec(json.NewCodec(), "application/json")
	//muxServer.RegisterService(new(GorillaStringService), "")

	router.Use(middlewares.CreateCorsMiddleware(server.ServerConfig.AllowedHosts))
	router.Use(middlewares.CreateLoggingMiddleware(os.Stdout, "Word Trainer"))
	router.Use(middlewares.CreatePanicRecoveryMiddleware())

	router.HandleFunc("/users/", server.SignUpAPI).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/session/", server.IsLogin).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/session/", server.Logout).Methods(http.MethodPatch, http.MethodOptions)
	router.HandleFunc("/session/", server.LoginAPI).Methods(http.MethodPost, http.MethodOptions)

	server.Router = router

	return server, nil
}

func (server *AuthServer) Run() {
	port := server.ServerConfig.Port
	server.Logger.Logf("Running app on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, server.Router))
}
