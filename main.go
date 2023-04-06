package main

import (
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/thetnaingtn/tidy-url/handlers"
)

func main() {

	db, err := sqlx.Connect("postgres", "user=postgres dbname=tidyurl password=postgres sslmode=disable")
	if err != nil {
		log.Println(err)
	}

	router := handlers.InitializeRouter(db)
	log.Fatal(http.ListenAndServe(":5000", router))
}
