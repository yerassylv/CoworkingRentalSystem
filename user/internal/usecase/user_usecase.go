package usecase

import (
	"context"
	"errors"
	"time"
	"user/internal/entity"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidCredentials = errors.New("invalid email or password")
	jwtSecret             = []byte("your_secret_key") // Замени на настоящий секрет
)

type UserRepository interface {
	GetUserByID(ctx context.Context, userID string) (*entity.User, error)
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
	CreateUser(ctx context.Context, user *entity.User) error
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

func (uc *UserUseCase) RegisterUser(ctx context.Context, fullName, email, phone, password string) (*entity.User, error) {
	passHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &entity.User{
		UserID:       uuid.NewString(),
		FullName:     fullName,
		Email:        email,
		Phone:        phone,
		PasswordHash: string(passHash),
	}

	err = uc.repo.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (uc *UserUseCase) LoginUser(ctx context.Context, email, password string) (string, *entity.User, error) {
	user, err := uc.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return "", nil, err
	}
	if user == nil {
		return "", nil, ErrInvalidCredentials
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return "", nil, ErrInvalidCredentials
	}

	// Генерация JWT токена
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  user.UserID,
		"email":    user.Email,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // токен на 24 часа
		"issuedAt": time.Now().Unix(),
	})

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", nil, err
	}

	return tokenString, user, nil
}
