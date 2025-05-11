type BookingUseCase struct {
    repo      domain.BookingRepository
    publisher domain.EventPublisher
}

func NewBookingUseCase(repo domain.BookingRepository, publisher domain.EventPublisher) *BookingUseCase {
    return &BookingUseCase{repo: repo, publisher: publisher}
}

func (uc *BookingUseCase) Create(...) error {
    booking := &domain.Booking{...}
    err := uc.repo.Create(booking)
    if err != nil {
        return err
    }
    return uc.publisher.PublishBookingCreated(booking)
}
