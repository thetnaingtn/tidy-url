package handlers

import (
	"github.com/jmoiron/sqlx"
	"github.com/julienschmidt/httprouter"
	"github.com/thetnaingtn/tidy-url/core"
)

func InitializeRouter(db *sqlx.DB) *httprouter.Router {
	router := httprouter.New()

	base := Handlers{
		core: core.NewCore(db),
	}

	router.GET("/:id", base.Expand)
	router.POST("/tidy", base.Tidy)

	return router
}
