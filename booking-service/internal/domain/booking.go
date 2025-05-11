package domain

import "time"

type Booking struct {
	ID        string
	UserID    string
	SpaceID   string
	StartTime time.Time
	EndTime   time.Time
	CreatedAt time.Time
}

type BookingRepository interface {
	Create(*Booking) error
	GetByID(id string) (*Booking, error)
	ListByUser(userID string) ([]*Booking, error)
}
