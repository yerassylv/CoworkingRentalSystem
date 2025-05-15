package usecase

import (
	"context"
	"space/internal/entity"
)

type SpaceRepository interface {
	GetSpaceByID(ctx context.Context, spaceID string) (*entity.Space, error)
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
