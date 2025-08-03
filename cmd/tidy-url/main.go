package main

import (
	"os"

	_ "github.com/lib/pq"
	"github.com/thetnaingtn/tidy-url/internal/config"
	"github.com/thetnaingtn/tidy-url/server"
	"github.com/thetnaingtn/tidy-url/store"
)

func main() {

	// uiFS, err := fs.Sub(UI, "ui/dist")
	// if err != nil {
	// 	log.Println(err)
	// }

	// db, err := sqlx.Connect("postgres", os.Getenv("DB_CONNECTION_URL"))
	// if err != nil {
	// 	log.Println("Can't connect to DB", err)
	// }

	// router := handlers.InitializeRouter(db, uiFS)
	// log.Fatal(http.ListenAndServe(os.Getenv("PORT"), router))

	config := &config.Config{
		Addr:    getEnv("ADDR", ""),
		DSN:     getEnv("DB_CONNECTION_URL", "postgres://user:password@localhost:5432/tidyurl?sslmode=disable"),
		BaseURL: getEnv("BASE_URL", ""),
		Port:    getEnv("PORT", "8080"),
	}

	store, err := store.NewStore(config)
	if err != nil {
		panic(err)
	}

	if err := server.NewServer(store, config).Start(); err != nil {
		panic(err)
	}

}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
