package server

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"github.com/user/2019_1_newTeam2/config"
	"github.com/user/2019_1_newTeam2/database"
	"github.com/user/2019_1_newTeam2/filesystem"
)

type Server struct {
	Router       *mux.Router
	DB           database.DBInterface
	ServerConfig *config.Config
}

func NewServer(pathToConfig string) (*Server, error) {
	server := new(Server)

	newConfig, err := config.NewConfig(pathToConfig)
	if err != nil {
		return nil, err
	}
	server.ServerConfig = newConfig
	newDB, err := database.NewDataBase()
	if err != nil {
		return nil, err
	}
	server.DB = newDB

	err = filesystem.CreateDir(filepath.Join(server.ServerConfig.UploadPath, server.ServerConfig.AvatarsPath))
	if err != nil {
		return nil, err
	}
	router := mux.NewRouter()

	router.HandleFunc("/users", server.UsersPaginate).Queries("rows", "{rows}", "page", "{page}").Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/users/", server.GetUser).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/users/", server.UpdateUser).Methods(http.MethodPatch, http.MethodOptions)
	router.HandleFunc("/users/", server.DeleteUser).Methods(http.MethodDelete, http.MethodOptions)
	router.HandleFunc("/users/", server.SignUpAPI).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/session/", server.IsLogin).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/session/", server.Logout).Methods(http.MethodPatch, http.MethodOptions)
	router.HandleFunc("/session/", server.LoginAPI).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/upload/{[0-9]+}", server.UploadAvatar).Methods(http.MethodPost, http.MethodOptions)

	router.PathPrefix("/files/{.+\\..+$}").Handler(http.StripPrefix("/files/", http.FileServer(http.Dir(server.ServerConfig.UploadPath))))

	server.Router = router

	return server, nil
}

func (server *Server) Run() {
	port := os.Getenv("PORT")
	if port == "" {
		port = server.ServerConfig.Port
	} // change for getting from config

	c := cors.New(cors.Options{
		AllowedHeaders:     []string{"Access-Control-Allow-Origin", "Charset", "Content-Type"},
		AllowedOrigins:     []string{"http://localhost:3000", "https://thawing-gorge-14317.herokuapp.com", "http://localhost:8090", "https://warm-shelf-71623.herokuapp.com"},
		AllowCredentials:   true,
		AllowedMethods:     []string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "DELETE", "PATCH"},
		OptionsPassthrough: true,
		Debug:              true,
	})

	handler := c.Handler(server.Router)
	http.ListenAndServe(":"+port, handler)
}
