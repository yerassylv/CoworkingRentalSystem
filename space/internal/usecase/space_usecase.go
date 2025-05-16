package usecase

import (
	"context"
	"space/internal/entity"

	"github.com/google/uuid"
)

type SpaceRepository interface {
	GetSpaceByID(ctx context.Context, spaceID string) (*entity.Space, error)
	CreateSpace(ctx context.Context, space *entity.Space) error
	UpdateSpace(ctx context.Context, space *entity.Space) error
	DeleteSpace(ctx context.Context, spaceID string) error
	ListSpaces(ctx context.Context) ([]*entity.Space, error)
}

type SpaceCache interface {
	GetSpace(ctx context.Context, spaceID string) (*entity.Space, error)
	SetSpace(ctx context.Context, space *entity.Space) error
	InvalidateSpace(ctx context.Context, spaceID string) error
}

type SpaceUseCase struct {
	repo  SpaceRepository
	cache SpaceCache
}

func NewSpaceUseCase(r SpaceRepository, c SpaceCache) *SpaceUseCase {
	return &SpaceUseCase{repo: r, cache: c}
}

func (uc *SpaceUseCase) GetSpaceByID(ctx context.Context, spaceID string) (*entity.Space, error) {
	space, err := uc.cache.GetSpace(ctx, spaceID)
	if err != nil {
		return nil, err
	}
	if space != nil {
		return space, nil
	}

	space, err = uc.repo.GetSpaceByID(ctx, spaceID)
	if err != nil {
		return nil, err
	}
	if space != nil {
		_ = uc.cache.SetSpace(ctx, space)
	}
	return space, nil
}

func (uc *SpaceUseCase) CreateSpace(ctx context.Context, name, location string, capacity int32, available bool) (*entity.Space, error) {
	space := &entity.Space{
		SpaceID:   uuid.NewString(),
		Name:      name,
		Location:  location,
		Capacity:  capacity,
		Available: available,
	}

	err := uc.repo.CreateSpace(ctx, space)
	if err != nil {
		return nil, err
	}
	_ = uc.cache.SetSpace(ctx, space)
	return space, nil
}

func (uc *SpaceUseCase) UpdateSpace(ctx context.Context, space *entity.Space) error {
	err := uc.repo.UpdateSpace(ctx, space)
	if err != nil {
		return err
	}
	_ = uc.cache.SetSpace(ctx, space)
	return nil
}

func (uc *SpaceUseCase) DeleteSpace(ctx context.Context, spaceID string) error {
	err := uc.repo.DeleteSpace(ctx, spaceID)
	if err != nil {
		return err
	}
	_ = uc.cache.InvalidateSpace(ctx, spaceID)
	return nil
}

func (uc *SpaceUseCase) ListSpaces(ctx context.Context) ([]*entity.Space, error) {
	return uc.repo.ListSpaces(ctx)
}
