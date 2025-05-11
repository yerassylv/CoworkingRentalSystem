package nats

import (
	"encoding/json"
	"log"

	"github.com/nats-io/nats.go"
)

type BookingEvent struct {
	SpaceID string `json:"space_id"`
}

func SubscribeToBookingCreated(nc *nats.Conn, handler func(string) error) {
	_, err := nc.Subscribe("booking.created", func(m *nats.Msg) {
		var event BookingEvent
		if err := json.Unmarshal(m.Data, &event); err != nil {
			log.Println("Error parsing event:", err)
			return
		}

		if err := handler(event.SpaceID); err != nil {
			log.Println("Error handling booking:", err)
		} else {
			log.Println("Handled booking.created for space:", event.SpaceID)
		}
	})
	if err != nil {
		log.Fatal("NATS subscription failed:", err)
	}
}
