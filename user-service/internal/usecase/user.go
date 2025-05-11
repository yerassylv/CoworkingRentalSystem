package usecase

import "user-service/internal/domain"

type UserUseCase struct {
	repo domain.UserRepository
}

func NewUserUseCase(repo domain.UserRepository) *UserUseCase {
	return &UserUseCase{repo: repo}
}

func (uc *UserUseCase) Register(u *domain.User) error {
	return uc.repo.Register(u)
}

func (uc *UserUseCase) GetByID(id string) (*domain.User, error) {
	return uc.repo.GetByID(id)
}

func (uc *UserUseCase) List() ([]*domain.User, error) {
	return uc.repo.ListAll()
}
