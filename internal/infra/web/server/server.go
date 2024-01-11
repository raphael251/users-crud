package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server struct {
	Router     chi.Router
	Handlers   map[string]http.HandlerFunc
	ServerPort string
}

func NewServer(serverPort string) *Server {
	return &Server{
		Router:     chi.NewRouter(),
		Handlers:   make(map[string]http.HandlerFunc),
		ServerPort: serverPort,
	}
}

func (s *Server) AddHandler(path string, handler http.HandlerFunc) {
	s.Handlers[path] = handler
}

func (s *Server) registerMiddlewares() {
	s.Router.Use(middleware.Logger)
	s.Router.Use(middleware.Recoverer)
}

func (s *Server) registerHandlers() {
	for path, handler := range s.Handlers {
		s.Router.Handle(path, handler)
	}
}

func (s *Server) Start() {
	s.registerMiddlewares()
	s.registerHandlers()

	log.Printf("Server listening on port %s", s.ServerPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", s.ServerPort), s.Router))
}
