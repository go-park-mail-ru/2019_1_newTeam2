package server

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/user/2019_1_newTeam2/pkg/middlewares"

	//"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	//"github.com/rs/cors"

	"github.com/user/2019_1_newTeam2/filesystem"
	"github.com/user/2019_1_newTeam2/pkg/config"
	"github.com/user/2019_1_newTeam2/pkg/logger"
	"github.com/user/2019_1_newTeam2/storage"
	"github.com/user/2019_1_newTeam2/storage/interfaces"
)

type Server struct {
	Router       *mux.Router
	DB           interfaces.DBInterface
	ServerConfig *config.Config
	Logger       logger.LoggerInterface
	CookieField  string
}

func NewServer(pathToConfig string) (*Server, error) {
	server := new(Server)

	logger := new(logger.GoLogger)
	logger.SetOutput(os.Stderr)
	logger.SetPrefix("LOG: ")
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

	err = filesystem.CreateDir(filepath.Join(server.ServerConfig.UploadPath, server.ServerConfig.AvatarsPath))
	if err != nil {
		return nil, err
	}
	router := mux.NewRouter()

	router.Use(middlewares.CreateCorsMiddleware(server.ServerConfig.AllowedHosts))
	router.Use(middlewares.CreateLoggingMiddleware(os.Stdout, "Word Trainer"))
	router.Use(middlewares.CreatePanicRecoveryMiddleware())

	needLogin := router.PathPrefix("/").Subrouter()
	needLogin.Use(middlewares.CreateCheckAuthMiddleware([]byte(server.ServerConfig.Secret), server.CookieField, IsLogined))
	// "need login"
	needLogin.HandleFunc("/users/", server.GetUser).Methods(http.MethodGet, http.MethodOptions)
	needLogin.HandleFunc("/users/", server.UpdateUser).Methods(http.MethodPut, http.MethodOptions)
	needLogin.HandleFunc("/users/", server.DeleteUser).Methods(http.MethodDelete, http.MethodOptions)
	needLogin.HandleFunc("/avatars/", server.UploadAvatar).Methods(http.MethodPost, http.MethodOptions)
	needLogin.HandleFunc("/languages/", server.GetLangs).Methods(http.MethodGet, http.MethodOptions)

	needLogin.HandleFunc("/dictionary", server.DictsPaginate).Queries("rows", "{rows}", "page", "{page}").Methods(http.MethodGet, http.MethodOptions)
	needLogin.HandleFunc("/dictionary/{id:[0-9]+}", server.GetDictionaryById).Methods(http.MethodGet, http.MethodOptions)
	needLogin.HandleFunc("/dictionary/", server.LoginAPI).Methods(http.MethodDelete, http.MethodOptions)
	needLogin.HandleFunc("/dictionary/", server.LoginAPI).Methods(http.MethodPost, http.MethodOptions)

	// mb to need login?
	needLogin.HandleFunc("/cards{dictId:[0-9]+}", server.CardsPaginate).Queries("rows", "{rows}", "page", "{page}").Methods(http.MethodGet, http.MethodOptions)
	needLogin.HandleFunc("/card/{id:[0-9]+}", server.GetCardById).Methods(http.MethodGet, http.MethodOptions)
	needLogin.HandleFunc("/card/", server.LoginAPI).Methods(http.MethodPut, http.MethodOptions)
	needLogin.HandleFunc("/card/", server.LoginAPI).Methods(http.MethodDelete, http.MethodOptions)
	needLogin.HandleFunc("/card/", server.LoginAPI).Methods(http.MethodPost, http.MethodOptions)
	//  end "need login"

	router.HandleFunc("/users", server.UsersPaginate).Queries("rows", "{rows}", "page", "{page}").Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/users/", server.SignUpAPI).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/session/", server.IsLogin).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/session/", server.Logout).Methods(http.MethodPatch, http.MethodOptions)
	router.HandleFunc("/session/", server.LoginAPI).Methods(http.MethodPost, http.MethodOptions)

	router.PathPrefix("/files/{.+\\..+$}").Handler(http.StripPrefix("/files/", http.FileServer(http.Dir(server.ServerConfig.UploadPath)))).Methods(http.MethodOptions, http.MethodGet)

	server.Router = router

	return server, nil
}

func (server *Server) Run() {
	port := os.Getenv("PORT")
	if port == "" {
		port = server.ServerConfig.Port
	}
	server.Logger.Logf("Running app on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, server.Router))
}
