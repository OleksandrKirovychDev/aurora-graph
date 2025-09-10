package internal

import (
	"aurora-graph/account/models"
	"context"
	"database/sql"
)

type Repository interface {
	Close()
	CreateAccount(ctx context.Context, account models.Account) (*models.Account, error)
	GetAccountByEmail(ctx context.Context, email string) (*models.Account, error)
	GetAccountById(ctx context.Context, id uint64) (*models.Account, error)
	ListAccounts(ctx context.Context, skip uint64, take uint64) ([]*models.Account, error)
}

type postgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) (Repository, error) {
	err := db.Ping()

	if err != nil {
		return nil, err
	}

	return &postgresRepository{db}, nil
}

func (repository *postgresRepository) Close() {}

func (repository *postgresRepository) CreateAccount(ctx context.Context, account models.Account) (*models.Account, error) {}

func (repository *postgresRepository) GetAccountByEmail(ctx context.Context, email string) (*models.Account, error) {}

func (repository *postgresRepository) GetAccountById(ctx context.Context, id uint64) (*models.Account, error) {}

func (repository *postgresRepository) ListAccounts(ctx context.Context, skip uint64, take uint64) ([]*models.Account, error) {}
