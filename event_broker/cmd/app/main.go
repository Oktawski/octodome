package main

import (
	"log"
	"net/http"
	"os"

	infra "octodome.com/eventbroker/infrastructure"
	"octodome.com/eventbroker/presentation"

	"github.com/go-chi/chi/v5"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "host=localhost user=sa password=pass123 dbname=octodome_event_db port=5433 sslmode=disable"
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	infra.Migrate(db)

	r := chi.NewRouter()
	presentation.Initialize(r, db)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8990"
	}

	log.Printf("Event broker service starting on port %s", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
