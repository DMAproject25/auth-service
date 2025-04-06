package usecase

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/DMAproject25/auth-service/internal/entity"
	"github.com/DMAproject25/auth-service/pkg/jwt"
)

type UserUseCase struct {
	repo      UserRepo
	secret    string
	token_TTL time.Duration
}

var _ Users = (*UserUseCase)(nil)

func NewUserUseCase(repo UserRepo, secret string, token_TTL time.Duration) *UserUseCase {
	return &UserUseCase{
		repo,
		secret,
		token_TTL,
	}
}

func (t *UserUseCase) SaveUser(ctx context.Context, email string) error {
	err := t.repo.SaveUser(ctx, email)

	if err != nil {
		return fmt.Errorf("can't save user: %w", err)
	}

	return nil
}

func (t *UserUseCase) UserByID(ctx context.Context, id uint64) (*entity.User, error) {
	user, err := t.repo.GetUserByID(ctx, id)

	if errors.Is(err, entity.ErrUserNotFound) {
		return nil, err
	}
	if err != nil {
		return nil, fmt.Errorf("can't get user by id: %w", err)
	}

	return user, nil
}

func (t *UserUseCase) Users(ctx context.Context) ([]entity.User, error) {
	users, err := t.repo.GetAllUsers(ctx)

	if err != nil {
		return nil, fmt.Errorf("can't get all users: %w", err)
	}

	return users, nil
}

func (t *UserUseCase) UserByEmail(ctx context.Context, email string) (*entity.User, error) {
	user, err := t.repo.GetUserByEmail(ctx, email)
	if errors.Is(err, entity.ErrUserNotFound) {
		return nil, err
	}
	if err != nil {
		return nil, fmt.Errorf("can't get user by email: %w", err)
	}

	return user, nil
}

func (u *UserUseCase) GenerateToken(userId uint64) (string, error) {
	payload := map[string]any{
		"uid": userId,
	}

	token, err := jwt.NewToken(payload, u.secret, u.token_TTL)
	if err != nil {
		return "", fmt.Errorf("can't generate token: %w", err)
	}

	return token, nil
}
