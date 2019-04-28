package authorization

import (
	"log"
	"net"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/user/2019_1_newTeam2/pkg/config"
	"github.com/user/2019_1_newTeam2/pkg/logger"
	"github.com/user/2019_1_newTeam2/pkg/middlewares"
	"github.com/user/2019_1_newTeam2/storage"
	"github.com/user/2019_1_newTeam2/storage/interfaces"
	"google.golang.org/grpc"
)

type AuthServer struct {
	Router       *mux.Router
	DB           interfaces.DBInterface
	ServerConfig *config.Config
	Logger       logger.LoggerInterface
	CookieField  string
	rpcServer	 *grpc.Server
	AuthClient	 AuthCheckerClient
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

	server.CookieField = "session_id"

	server.ServerConfig = newConfig
	newDB, err := storage.NewDataBase(server.ServerConfig.DBUser, server.ServerConfig.DBPassUser)
	if err != nil {
		server.Logger.Log(err)
		return nil, err
	}
	server.DB = newDB

	router := mux.NewRouter()

	server.rpcServer = grpc.NewServer()
	RegisterAuthCheckerServer(server.rpcServer, NewAuthManager())

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
	lis, err := net.Listen("tcp", ":8092")
	if err != nil {
		server.Logger.Logf("Can`t listen port %s", "8092")
		return
	}
	go server.rpcServer.Serve(lis)
	server.Logger.Logf("Running AuthMS(grps) on port %s", "8092")

	grcpAuthConn, err := grpc.Dial(
		"127.0.0.1:8092",
		grpc.WithInsecure(),
	)
	if err != nil {
		server.Logger.Log("Can`t connect ro grpc (auth ms)")
		return
	}

	defer grcpAuthConn.Close()
	server.AuthClient = NewAuthCheckerClient(grcpAuthConn)

	port := server.ServerConfig.Port
	server.Logger.Logf("Running app on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, server.Router))
}
