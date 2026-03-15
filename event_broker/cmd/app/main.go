package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"octodome.com/eventbroker/application"
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

	registry := infra.NewHandlerRegistry(db)
	dispatcher := infra.NewEventDispatcher()
	eventRepository := infra.NewEventRepository(db)
	stateManager := infra.NewStateManager(db)

	sweeper := infra.NewSweeper(eventRepository, registry, dispatcher)

	forward := application.NewForward(eventRepository, registry, dispatcher)
	updateState := application.NewUpdateState(stateManager)
	getEvent := application.NewGetEvent(eventRepository)
	registerHandler := application.NewRegisterHandler(registry)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// TODO: make cleaner goroutine that will mark processing events as failed after 5 minutes
	go func() {
		ticker := time.NewTicker(30 * time.Second)
		defer ticker.Stop()
		log.Printf("Sweeper started with interval %s", 30*time.Second)
		for {
			select {
			case <-ticker.C:
				if err := sweeper.Sweep(ctx); err != nil {
					log.Printf("Sweep failed: %v", err)
				}
			case <-ctx.Done():
				return
			}
		}
	}()

	r := chi.NewRouter()
	presentation.RegisterEventRoutes(r, forward, updateState, getEvent, registerHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8990"
	}

	log.Printf("Event broker service starting on port %s", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
