package domain

type EventPublisher interface {
	PublishBookingCreated(booking *Booking) error
}
