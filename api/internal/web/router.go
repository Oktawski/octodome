package web

import (
	"log"
	"net/http"
	"os"

	authmod "octodome.com/api/internal/auth"
	equipmentmod "octodome.com/api/internal/equipment/mod"
	usermod "octodome.com/api/internal/user"

	"github.com/go-chi/chi/v5"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func StartServer() {
	r := chi.NewRouter()

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "host=localhost user=sa password=pass123 dbname=octodome_db port=5432 sslmode=disable"
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	r.Route("/api/v1", func(r chi.Router) {
		authmod.Initialize(r, db)
		usermod.Initialize(r, db)
		equipmentmod.Initialize(r, db)
	})

	http.ListenAndServe(":8989", r)
}
