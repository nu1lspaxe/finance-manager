package users

import (
	"context"
	"finance_manager/configs"
	"finance_manager/pkg/postgres/sqlc"

	"github.com/jackc/pgx/v5/pgxpool"
)

type IUserService interface {
	CreateUser(ctx context.Context, params sqlc.CreateUserParams) (*sqlc.User, error)
}

type UserService struct {
	pool    *pgxpool.Pool
	queries *sqlc.Queries
}

func NewUserService(pool *pgxpool.Pool) IUserService {
	return &UserService{
		pool:    pool,
		queries: sqlc.New(pool),
	}
}

func (s *UserService) CreateUser(ctx context.Context, params sqlc.CreateUserParams) (*sqlc.User, error) {
	if params.Username == "" {
		return nil, configs.NewManagerError(configs.ErrInvalidValue, "username is required")
	}
	if params.Email == "" {
		return nil, configs.NewManagerError(configs.ErrInvalidValue, "email is required")
	}

	exists, err := s.queries.CheckUserEmailExists(ctx, params.Email)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, configs.NewManagerError(configs.ErrUserExists, "user already exists")
	}

	user, err := s.queries.CreateUser(ctx, params)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
