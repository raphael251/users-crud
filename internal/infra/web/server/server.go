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
	Routers    map[string]RouterFunc
	ServerPort string
}

type RouterFunc func(r chi.Router)

func NewServer(serverPort string) *Server {
	return &Server{
		Router:     chi.NewRouter(),
		Routers:    make(map[string]RouterFunc),
		ServerPort: serverPort,
	}
}

func (s *Server) AddRouter(path string, routerFunc RouterFunc) {
	s.Routers[path] = routerFunc
}

func (s *Server) registerMiddlewares() {
	s.Router.Use(middleware.Logger)
	s.Router.Use(middleware.Recoverer)
}

func (s *Server) registerRouters() {
	for path, router := range s.Routers {
		s.Router.Route(path, router)
	}
}

func (s *Server) Start() {
	s.registerMiddlewares()
	s.registerRouters()

	log.Printf("Server listening on port %s", s.ServerPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", s.ServerPort), s.Router))
}
