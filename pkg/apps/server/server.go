package server

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"

	// "github.com/user/2019_1_newTeam2/pkg/wshub"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"google.golang.org/grpc"

	"github.com/user/2019_1_newTeam2/pkg/apps/authorization"
	"github.com/user/2019_1_newTeam2/shared/pkg/middlewares"

	"github.com/gorilla/mux"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/user/2019_1_newTeam2/shared/filesystem"
	"github.com/user/2019_1_newTeam2/shared/pkg/config"
	"github.com/user/2019_1_newTeam2/shared/pkg/logger"
	"github.com/user/2019_1_newTeam2/shared/storage"
	"github.com/user/2019_1_newTeam2/shared/storage/interfaces"
)

type Server struct {
	Router       *mux.Router
	DB           interfaces.DBInterface
	ServerConfig *config.Config
	Logger       logger.LoggerInterface
	CookieField  string
	// Hub          wshub.IWSCommunicator
	AuthClient authorization.AuthCheckerClient
	svc        *s3.S3
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

	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String("ru-msk"),
		Credentials: credentials.NewStaticCredentials(newConfig.Akid, newConfig.CloudSecret, ""),
		Endpoint:    aws.String("http://hb.bizmrg.com"),
	}))
	server.svc = s3.New(sess)

	server.CookieField = "session_id"

	server.ServerConfig = newConfig
	newDB, err := storage.NewDataBase(server.ServerConfig.DBHost, server.ServerConfig.DBUser, server.ServerConfig.DBPassUser)
	if err != nil {
		return nil, err
	}
	server.DB = newDB

	err = filesystem.CreateDir(filepath.Join(server.ServerConfig.UploadPath, server.ServerConfig.AvatarsPath))
	if err != nil {
		return nil, err
	}

	r := mux.NewRouter()

	router := r.PathPrefix("/api/").Subrouter()
	router.Use(middlewares.CreateCorsMiddleware(server.ServerConfig.AllowedHosts))
	router.Use(middlewares.CreateLoggingMiddleware(os.Stdout, "Word Trainer"))
	router.Use(middlewares.CreatePanicRecoveryMiddleware())
	router.Use(server.metricsMiddleware)

	needLogin := router.PathPrefix("/").Subrouter()
	needLogin.Use(middlewares.CreateCheckAuthMiddleware([]byte(server.ServerConfig.Secret), server.CookieField, server.IsLogined))
	needLogin.HandleFunc("/users/", server.GetUser).Methods(http.MethodGet, http.MethodOptions)
	needLogin.HandleFunc("/users/", server.UpdateUser).Methods(http.MethodPut, http.MethodOptions)
	needLogin.HandleFunc("/users/", server.DeleteUser).Methods(http.MethodDelete, http.MethodOptions)
	needLogin.HandleFunc("/avatars/", server.UploadAvatar).Methods(http.MethodPost, http.MethodOptions)
	needLogin.HandleFunc("/languages/", server.GetLangs).Methods(http.MethodGet, http.MethodOptions)

	needLogin.HandleFunc("/dictionary", server.DictsPaginate).Queries("rows", "{rows}", "page", "{page}").Methods(http.MethodGet, http.MethodOptions)
	needLogin.HandleFunc("/dictionary/{id:[0-9]+}", server.GetDictionaryById).Methods(http.MethodGet, http.MethodOptions)
	needLogin.HandleFunc("/dictionary/", server.UpdateDictionaryAPI).Methods(http.MethodPut, http.MethodOptions)
	needLogin.HandleFunc("/dictionary/{id:[0-9]+}", server.DeleteDictionaryAPI).Methods(http.MethodDelete, http.MethodOptions)
	needLogin.HandleFunc("/dictionary/", server.CreateDictionaryAPI).Methods(http.MethodPost, http.MethodOptions)

	needLogin.HandleFunc("/dictionary/{id:[0-9]+}", server.BorrowDictById).Methods(http.MethodPatch)

	needLogin.HandleFunc("/cards", server.CardsPaginate).Queries("dict", "{dictId}", "rows", "{rows}", "page", "{page}").Methods(http.MethodGet, http.MethodOptions)
	needLogin.HandleFunc("/card/{id:[0-9]+}", server.GetCardById).Methods(http.MethodGet, http.MethodOptions)
	needLogin.HandleFunc("/card/", server.DeleteCardInDictionary).Methods(http.MethodDelete, http.MethodOptions)
	needLogin.HandleFunc("/card/", server.CreateCardInDictionary).Methods(http.MethodPost, http.MethodOptions)
	needLogin.HandleFunc("/cards/", server.UploadWordsFileAPI).Methods(http.MethodPost, http.MethodOptions)

	needLogin.HandleFunc("/single", server.GetSingleGame).Queries("dict", "{dictId}", "words", "{wordsNum}").Methods(http.MethodGet, http.MethodOptions)
	needLogin.HandleFunc("/single", server.SetGameResults).Methods(http.MethodPost, http.MethodOptions)

	// set needLogin in future, when front is ready
	// needLogin.HandleFunc("/subscribe/", server.WSSubscribe).Methods(http.MethodGet)

	router.HandleFunc("/users", server.UsersPaginate).Queries("rows", "{rows}", "page", "{page}").Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/users/", server.SignUpAPI).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/session/", server.IsLogin).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/session/", server.Logout).Methods(http.MethodPatch, http.MethodOptions)
	router.HandleFunc("/session/", server.LoginAPI).Methods(http.MethodPost, http.MethodOptions)

	r.Handle("/metrics", promhttp.Handler())

	router.PathPrefix("/files/{.+\\..+$}").Handler(http.StripPrefix("/api/files/", http.FileServer(http.Dir(server.ServerConfig.UploadPath)))).Methods(http.MethodOptions, http.MethodGet)

	server.Router = router
	server.Router = r

	return server, nil
}

func (server *Server) Run() {
	port := os.Getenv("PORT")
	if port == "" {
		port = server.ServerConfig.Port
	}

	grcpAuthConn, err := grpc.Dial(
		server.ServerConfig.AuthHost+":"+server.ServerConfig.AuthPort,
		grpc.WithInsecure(),
	)
	if err != nil {
		server.Logger.Log("Can`t connect ro grpc (auth ms)")
	}
	defer grcpAuthConn.Close()

	server.AuthClient = authorization.NewAuthCheckerClient(grcpAuthConn)

	prometheus.MustRegister(ApiMetrics)
	server.Logger.Logf("Running app on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, server.Router))
}
