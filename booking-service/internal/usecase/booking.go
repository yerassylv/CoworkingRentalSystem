package usecase

import (
	"booking-service/internal/domain"
	"time"
)

type BookingUseCase struct {
	repo      domain.BookingRepository
	publisher domain.EventPublisher
}

func NewBookingUseCase(repo domain.BookingRepository, publisher domain.EventPublisher) *BookingUseCase {
	return &BookingUseCase{repo: repo, publisher: publisher}
}

func (uc *BookingUseCase) Create(userID, spaceID string, start, end time.Time) error {
	booking := &domain.Booking{
		UserID:    userID,
		SpaceID:   spaceID,
		StartTime: start,
		EndTime:   end,
		CreatedAt: time.Now(),
	}

	if err := uc.repo.Create(booking); err != nil {
		return err
	}

	return uc.publisher.PublishBookingCreated(booking)
}
