package frontend

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/thetnaingtn/tidy-url/internal/config"
	"github.com/thetnaingtn/tidy-url/store"
)

//go:embed dist/*
var embedFiles embed.FS

type FrontendService struct {
	Store  *store.Store
	Config *config.Config
}

func NewFrontendService(store *store.Store, config *config.Config) *FrontendService {
	return &FrontendService{
		Store:  store,
		Config: config,
	}
}

func (*FrontendService) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	distFS, err := fs.Sub(embedFiles, "dist")
	if err != nil {
		panic(err)
	}

	http.FileServer(http.FS(distFS)).ServeHTTP(w, r)
}
