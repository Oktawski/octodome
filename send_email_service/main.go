package main

import (
	"log"
	"net/http"
	"os"

	"octodome.com/shared/events"
)

func main() {
	brokerURL := os.Getenv("EVENT_BROKER_URL")
	serviceURL := os.Getenv("SERVICE_URL")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	eventClient := events.NewClient(brokerURL)

	if err := eventClient.RegisterHandler(
		"send_email_service",
		events.UserRegistered{}.GetEventType(),
		serviceURL+"/events/user-registered",
	); err != nil {
		log.Fatalf("Failed to register handler with event broker: %v", err)
	}
	log.Printf("Registered handler for %s at %s/events", events.UserRegistered{}.GetEventType(), serviceURL)

	if err := eventClient.RegisterHandler(
		"send_email_service",
		events.MagicCodeRequested{}.GetEventType(),
		serviceURL+"/events/magic-code-requested",
	); err != nil {
		log.Fatalf("Failed to register handler with event broker: %v", err)
	}
	log.Printf("Registered handler for %s at %s/events", events.MagicCodeRequested{}.GetEventType(), serviceURL)

	router := newRouter(*eventClient)

	log.Printf("Starting registration service on :%s", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
