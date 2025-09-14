package internal

import (
	"aurora-graph/account/models"
	"aurora-graph/pkg/auth"
	"aurora-graph/pkg/crypt"
	"context"
	"errors"
)

type Service interface {
	Register(ctx context.Context, name, email, password string) (string, error)
	Login(ctx context.Context, email, password string) (string, error)
	GetAccount(ctx context.Context, id uint64) (*models.Account, error)
	GetAccounts(ctx context.Context, skip, take uint64) ([]*models.Account, error)
}

type accountService struct {
	repository Repository
}

func NewAccountService(r Repository) Service {
	return &accountService{repository: r}
}

func (service *accountService) Register(ctx context.Context, name string, email string, password string) (string, error) {
	res, _ := service.repository.GetAccountByEmail(ctx, email)
	if res != nil {
		return "", errors.New("account with this email already exists")
	}

	hashedPassword, err := crypt.HashPassword(password)
	if err != nil {
		return "", err
	}

	acc := models.Account{
		Email: email,
		Name: name,
		Password: hashedPassword,
	}

	createdAccount, err := service.repository.CreateAccount(ctx, acc)
	if err != nil {
		return "", err
	}

	token, err := auth.GenerateToken(createdAccount.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (service *accountService) Login(ctx context.Context, email string, password string) (string, error) {
	account, err := service.repository.GetAccountByEmail(ctx, email)
	if err != nil {
		return "", err
	}

	err = crypt.VerifyPassword(password, account.Password)
	if err != nil {
		return "", err
	}

	token, err := auth.GenerateToken(account.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (service *accountService) GetAccount(ctx context.Context, id uint64) (*models.Account, error) {
	return service.repository.GetAccountById(ctx, id)
}

func (service *accountService) GetAccounts(ctx context.Context, skip uint64, take uint64) ([]*models.Account, error) {
	return service.repository.ListAccounts(ctx, skip, take)
}
