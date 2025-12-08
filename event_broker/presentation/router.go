package presentation

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
	infra "octodome.com/eventbroker/infrastructure"
)

func Initialize(r chi.Router, db *gorm.DB) {
	publisher := infra.NewEventPublisher(db)
	consumer := infra.NewEventConsumer(db)
	controller := NewEventController(publisher, consumer)

	r.Route("/events", func(events chi.Router) {
		events.Post("/", controller.PublishEvent)
		events.Get("/{eventType}", controller.GetEvent)
		events.Put("/{id:[0-9]+}/pending", controller.MarkEventAsPending)
		events.Put("/{id:[0-9]+}/processing", controller.MarkEventAsProcessing)
		events.Put("/{id:[0-9]+}/processed", controller.MarkEventAsProcessed)
		events.Put("/{id:[0-9]+}/failed", controller.MarkEventAsFailed)
	})

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	})
}
