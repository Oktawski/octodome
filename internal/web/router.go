package web

import (
	"log"
	"net/http"
	authmod "octodome/internal/auth"
	equipmentmod "octodome/internal/equipment/mod"
	usermod "octodome/internal/user"

	"github.com/go-chi/chi/v5"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func StartServer() {
	r := chi.NewRouter()

	dsn := "host=localhost user=sa password=pass123 dbname=octodome_db port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	authmod.Initialize(r, db)
	usermod.Initialize(r, db)
	equipmentmod.Initialize(r, db)

	http.ListenAndServe(":8989", r)
}
