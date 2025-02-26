package usecase

import (
	"context"

	"github.com/DMAproject25/auth-service/internal/entity"
)

type (
	Todos interface {
		Todos(ctx context.Context) ([]entity.Todo, error)
		TodoByID(ctx context.Context, id uint64) (*entity.Todo, error)
		SaveTodo(ctx context.Context, task string) error
	}

	TodosRepo interface {
		GetAllTodos(ctx context.Context) ([]entity.Todo, error)
		GetTodoByID(ctx context.Context, id uint64) (*entity.Todo, error)
		SaveTodo(ctx context.Context, task string) error
	}
)

type (
	Users interface {
		Users(ctx context.Context) ([]entity.User, error)
		UserByID(ctx context.Context, id uint64) (*entity.User, error)
		SaveUser(ctx context.Context, email string) error
	}

	UserRepo interface {
		GetAllUsers(ctx context.Context) ([]entity.User, error)
		GetUserByID(ctx context.Context, id uint64) (*entity.User, error)
		SaveUser(ctx context.Context, email string) error
	}
)
