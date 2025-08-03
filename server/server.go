package server

import (
	"github.com/thetnaingtn/tidy-url/internal/config"
	"github.com/thetnaingtn/tidy-url/store"
)

type Server struct {
	Store  *store.Store
	Config *config.Config
}

func NewServer(store *store.Store, config *config.Config) *Server {
	return &Server{
		Store:  store,
		Config: config,
	}
}

func (s *Server) Start() error {
	return nil
}
