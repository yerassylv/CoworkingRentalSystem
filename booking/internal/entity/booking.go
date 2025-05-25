package entity

type Booking struct {
	BookingID string `bson:"booking_id" json:"booking_id"`
	UserID    string `bson:"user_id" json:"user_id"`
	SpaceID   string `bson:"space_id" json:"space_id"`
	Date      string `bson:"date" json:"date"`
	Email     string `bson:"email" json:"email"`
	Status    string `bson:"status,omitempty" json:"status,omitempty"`

	// Optional, can be used for confirmation emails
}
