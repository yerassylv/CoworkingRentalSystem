package usecase

import "space-service/internal/domain"

type SpaceUseCase struct {
	repo domain.SpaceRepository
}

func NewSpaceUseCase(repo domain.SpaceRepository) *SpaceUseCase {
	return &SpaceUseCase{repo: repo}
}

func (uc *SpaceUseCase) HandleBooking(spaceID string) error {
	return uc.repo.DecreaseAvailability(spaceID)
}
