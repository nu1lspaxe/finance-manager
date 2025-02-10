package users

import (
	"context"
	"finance_manager/configs"
	"finance_manager/pkg/postgres/sqlc"
)

type IUserService interface {
	CreateUser(ctx context.Context, params sqlc.CreateUserParams) (*sqlc.User, error)
}

type UserService struct {
	queries *sqlc.Queries
}

func NewUserService(queries *sqlc.Queries) IUserService {
	return &UserService{
		queries: queries,
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
