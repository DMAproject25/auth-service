package repo

import (
	"context"
	"fmt"

	"github.com/DMAproject25/auth-service/internal/entity"
	"github.com/DMAproject25/auth-service/pkg/postgres"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
)

const (
	_defaultListCapUser = 64
)

type UserRepo struct {
	*postgres.Postgres
}

func NewUserRepo(pg *postgres.Postgres) *UserRepo {
	return &UserRepo{
		pg,
	}
}

func (t *UserRepo) GetAllUsers(ctx context.Context) ([]entity.User, error) {
	query, _, err := t.Builder.Select("id", "email").From("users").ToSql()
	if err != nil {
		return []entity.User{}, fmt.Errorf("can't create sql query: %w", err)
	}

	rows, err := t.Pool.Query(ctx, query)
	if err != nil {
		return []entity.User{}, fmt.Errorf("can't query request: %w", err)
	}
	defer rows.Close()

	users := make([]entity.User, 0, _defaultListCapUser)

	for rows.Next() {
		user := entity.User{}

		err := rows.Scan(&user.ID, &user.Email)
		if err != nil {
			return []entity.User{}, fmt.Errorf("can't scan rows: %w", err)
		}

		users = append(users, user)
	}

	return users, nil
}

func (t *UserRepo) GetUserByID(ctx context.Context, id uint64) (*entity.User, error) {
	query, args, err := t.Builder.
		Select("id", "email").
		From("users").
		Where(squirrel.Eq{"id": id}).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("can't create sql query: %w", err)
	}

	row := t.Pool.QueryRow(ctx, query, args...)

	result := entity.User{}

	err = row.Scan(&result.ID, &result.Email)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, entity.ErrTodoNotFound
		}

		return nil, fmt.Errorf("can't to scan row: %w", err)
	}

	return &result, nil
}

func (t *UserRepo) SaveUser(ctx context.Context, email string) error {
	query, args, err := t.Builder.
		Insert("todos").
		Columns("email").
		Values(email).
		ToSql()
	if err != nil {
		return fmt.Errorf("can't create sql query: %w", err)
	}

	if _, err = t.Pool.Exec(ctx, query, args...); err != nil {
		return fmt.Errorf("can't insert into table: %w", err)
	}

	return nil
}
