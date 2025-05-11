package nats

import (
	"booking-service/internal/domain"
	"encoding/json"

	"github.com/nats-io/nats.go"
)

type BookingPublisher struct {
	conn *nats.Conn
}

func NewBookingPublisher(conn *nats.Conn) domain.EventPublisher {
	return &BookingPublisher{conn: conn}
}

func (p *BookingPublisher) PublishBookingCreated(b *domain.Booking) error {
	data, err := json.Marshal(b)
	if err != nil {
		return err
	}
	return p.conn.Publish("booking.created", data)
}
