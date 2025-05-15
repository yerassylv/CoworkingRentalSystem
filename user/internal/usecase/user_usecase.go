package usecase

import (
	"context"
	"user/internal/entity"
)

type UserRepository interface {
	GetUserByID(ctx context.Context, userID string) (*entity.User, error)
}

type UserCache interface {
	GetUserProfile(ctx context.Context, userID string) (*entity.User, error)
	SetUserProfile(ctx context.Context, user *entity.User) error
	InvalidateUserProfile(ctx context.Context, userID string) error
}

type UserUseCase struct {
	repo  UserRepository
	cache UserCache
}

func NewUserUseCase(r UserRepository, c UserCache) *UserUseCase {
	return &UserUseCase{repo: r, cache: c}
}

func (uc *UserUseCase) GetUserProfile(ctx context.Context, userID string) (*entity.User, error) {
	user, err := uc.cache.GetUserProfile(ctx, userID)
	if err != nil {
		return nil, err
	}
	if user != nil {
		return user, nil
	}

	user, err = uc.repo.GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	if user != nil {
		_ = uc.cache.SetUserProfile(ctx, user) // Ignore cache error
	}
	return user, nil
}
