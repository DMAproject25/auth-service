package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/DMAproject25/auth-service/internal/entity"
)

type UserUseCase struct {
	repo UserRepo
}

var _ Users = (*UserUseCase)(nil)

func NewUserUseCase(repo UserRepo) *UserUseCase {
	return &UserUseCase{
		repo,
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
