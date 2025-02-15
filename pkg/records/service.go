package records

import (
	"context"
	"finance_manager/configs"
	"finance_manager/pkg/postgres/sqlc"

	"github.com/jackc/pgx/v5/pgxpool"
)

type IRecordService interface {
	CreateRecord(ctx context.Context, params sqlc.CreateRecordParams) (*sqlc.Record, error)
}

type RecordService struct {
	pool    *pgxpool.Pool
	queries *sqlc.Queries
}

func NewRecordService(pool *pgxpool.Pool) IRecordService {
	return &RecordService{
		pool:    pool,
		queries: sqlc.New(pool),
	}
}

func (s *RecordService) CreateRecord(ctx context.Context, params sqlc.CreateRecordParams) (*sqlc.Record, error) {
	exists, err := s.queries.CheckUserExists(ctx, params.UserID)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, configs.NewManagerError(configs.ErrUserNotFound, "user not found")
	}

	if params.Amount.Int.Int64() <= 0 {
		return nil, configs.NewManagerError(configs.ErrInvalidValue, "amount must be greater than 0")
	}

	record, err := s.queries.CreateRecord(ctx, params)
	if err != nil {
		return nil, err
	}
	return &record, nil
}
