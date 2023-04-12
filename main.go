package main

import (
	"io/fs"
	"log"
	"net/http"
	"os"

	"embed"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/thetnaingtn/tidy-url/handlers"
)

//go:embed ui/dist
var UI embed.FS

func main() {

	uiFS, err := fs.Sub(UI, "ui/dist")
	if err != nil {
		log.Println(err)
	}

	db, err := sqlx.Connect("postgres", os.Getenv("DB_CONNECTION_URL"))
	if err != nil {
		log.Println(err)
	}

	router := handlers.InitializeRouter(db, uiFS)
	log.Fatal(http.ListenAndServe(":5000", router))
}
