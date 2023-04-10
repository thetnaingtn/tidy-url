package handlers

import (
	"io/fs"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/julienschmidt/httprouter"
	"github.com/thetnaingtn/tidy-url/core"
)

type Handlers struct {
	core core.Core
	uiFS fs.FS
}

func InitializeRouter(db *sqlx.DB, uiFS fs.FS) *httprouter.Router {
	router := httprouter.New()

	base := Handlers{
		core: core.NewCore(db),
		uiFS: uiFS,
	}

	router.ServeFiles("/assets/*filepath", http.Dir("ui/dist/assets"))

	router.GET("/", base.StaticHandler)
	router.GET("/expand/:id", base.Expand)
	router.POST("/tidy", base.Tidy)

	return router
}
