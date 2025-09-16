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

func (repository *postgresRepository) Close() {
	repository.db.Close()
}

func (repository *postgresRepository) CreateAccount(ctx context.Context, account models.Account) (*models.Account, error) {
	query := `
		INSERT INTO accounts (email, name, password, created_at)
		VALUES ($1, $2, $3, NOW())
		RETURNING id, email, name, created_at
	`

	row := repository.db.QueryRowContext(ctx, query, account.Email, account.Name, account.Password)

	var newAccount models.Account

	err := row.Scan(&newAccount.ID, &newAccount.Email, &newAccount.Name, &newAccount.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &newAccount, nil
}

func (repository *postgresRepository) GetAccountByEmail(ctx context.Context, email string) (*models.Account, error) {
	query := `
		SELECT id, name, email, password, created_at
		FROM accounts
		WHERE email = $1
	`

	row := repository.db.QueryRowContext(ctx, query, email)

	var account models.Account

	err := row.Scan(&account.ID, &account.Name, &account.Email, &account.CreatedAt, &account.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &account, nil
}

func (repository *postgresRepository) GetAccountById(ctx context.Context, id uint64) (*models.Account, error) {
	query := `
		SELECT id, name, email, created_at, updated_at
		FROM accounts
		WHERE id = $1
	`

	row := repository.db.QueryRowContext(ctx, query, id)

	var account models.Account

	err := row.Scan(&account.ID, &account.Name, &account.Email, &account.CreatedAt, &account.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &account, nil
}

func (repository *postgresRepository) ListAccounts(ctx context.Context, skip uint64, take uint64) ([]*models.Account, error) {
	query := `
		SELECT id, name, email, created_at
		FROM accounts
		OFFSET $1 LIMIT $2
	`

	rows, err := repository.db.QueryContext(ctx, query, skip, take)
	if err != nil {
		return nil, err
	}

	var accounts []*models.Account

	for rows.Next() {
		var account models.Account

		err := rows.Scan(&account.ID, &account.Email, &account.Name, &account.CreatedAt)
		if err != nil {
			return nil, err
		}

		accounts = append(accounts, &account)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return accounts, nil
}
