package apiserver

import (
	"github.com/WelchDragon/http-rest-api.git/internal/app/store"
	"github.com/sirupsen/logrus"

	"github.com/gorilla/mux"

	"net/http"

	"io"
)

//APIServer ...
type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store  *store.Store
}

//New ...
func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

//Start ...
func (s *APIServer) Start() error {
	if err := s.configugerLogger(); err != nil {
		return err
	}
	s.configugerRouter()

	if err := s.configugerStore(); err != nil {
		return err
	}

	s.logger.Info("starting api server")

	return http.ListenAndServe(s.config.BindAddr, s.router)
}

//configurerLogger ...
func (s *APIServer) configugerLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)
	return nil
}

func (s *APIServer) configugerRouter() {
	s.router.HandleFunc("/hello", s.handleHello())
}

func (s *APIServer) configugerStore() error {
	st := store.New(s.config.Store)
	if err := st.Open(); err != nil {
		return err
	}
	s.store = st
	return nil
}

func (s *APIServer) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello")
	}
}
