package main

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"octodome.com/shared/events"
)

func main() {
	brokerURL := os.Getenv("EVENT_BROKER_URL")
	eventClient := events.NewClient(brokerURL)
	ticker := time.NewTicker(5 * time.Second)

	defer ticker.Stop()

	processEvent(*eventClient)
	for range ticker.C {
		processEvent(*eventClient)
	}
}

func processEvent(eventClient events.Client) {
	id, event, err := eventClient.GetEvent(events.UserRegistered{}.GetEventType())
	if err != nil {
		log.Default().Printf("Error getting event %s: %v", events.UserRegistered{}.GetEventType(), err)
		return
	}

	var userRegisteredPayload events.UserRegistered
	err = json.Unmarshal(event, &userRegisteredPayload)
	if err != nil {
		err := eventClient.MarkEventAsFailed(id)
		log.Default().Printf("Error unmarshalling event %s: %v", events.UserRegistered{}.GetEventType(), err)
		return
	}

	processed, err := sendEmail(userRegisteredPayload)

	if err != nil {
		log.Fatalf("Error processing event %s: %v", events.UserRegistered{}.GetEventType(), err)
		err := eventClient.MarkEventAsFailed(id)
		if err != nil {
			log.Default().Print(err.Error())
		}
	}

	if processed {
		err := eventClient.MarkEventAsProcessed(id)
		if err != nil {
			log.Default().Print(err.Error())
		}
	}
}

func sendEmail(userRegisteredPayload events.UserRegistered) (bool, error) {
	return true, nil
}
