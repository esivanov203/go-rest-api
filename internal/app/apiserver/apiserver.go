package apiserver

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

type ApiServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
}

func New(config *Config) *ApiServer {
	return &ApiServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (s *ApiServer) Run() error {
	if err := s.configureLogger(); err != nil {
		return err
	}
	s.configureRouter()
	s.logger.Info("Starting API server")

	return http.ListenAndServe(s.config.BindAddr, s.router)
}
