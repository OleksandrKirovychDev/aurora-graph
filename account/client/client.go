package client

import (
	"aurora-graph/account/models"
	"aurora-graph/account/proto/pb"
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	conn	*grpc.ClientConn
	service pb.AccountServiceClient
}

func NewClient(url string) (*Client, error) {
	conn, err := grpc.NewClient(url, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, err
	}

	C := pb.NewAccountServiceClient(conn)
	return &Client{conn, C}, nil
}

func (client *Client) Close() {
	err := client.conn.Close()

	if err != nil {
		log.Println(err)
	}
}

func (client *Client) Register(ctx context.Context, name, email, password string) (string, error) {
	response, err := client.service.Register(ctx, &pb.RegisterRequest{
		Email: email,
		Name: name,
		Password: password,
	})

	if err != nil {
		return "", err
	}

	return response.Value, nil
}

func (client *Client) Login(ctx context.Context, email, password string) (string, error) {
	response, err := client.service.Login(ctx, &pb.LoginRequest{
		Email: email,
		Password: password,
	})

	if err != nil {
		return "", err
	}

	return response.Value, nil
}

func (client *Client) GetAccount(ctx context.Context, Id uint64) (*models.Account, error) {
	response, err := client.service.GetAccount(ctx, &pb.GetAccountRequest{
		Id: Id,
	})

	if err != nil {
		return nil, err
	}

	var createdAt, updatedAt *time.Time

	if response.Account.CreatedAt != nil {
		t := response.Account.CreatedAt.AsTime()
		createdAt = &t
	}
	if response.Account.UpdatedAt != nil {
		t := response.Account.UpdatedAt.AsTime()
		updatedAt = &t
	}

	return &models.Account{
		ID: response.Account.Id,
		Email: response.Account.Email,
		Name: response.Account.Name,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}, nil
}

func (client *Client) GetAccounts(ctx context.Context, skip, take uint64) ([]models.Account, error) {
	response, err := client.service.GetAccounts(
		ctx,
		&pb.GetAccountsRequest{
			Take: take, Skip: skip,
		},
	)

	if err != nil {
		return nil, err
	}

	var accounts []models.Account

	for _, a := range response.Accounts {
		var createdAt, updatedAt *time.Time

		if a.CreatedAt != nil {
			t := a.CreatedAt.AsTime()
			createdAt = &t
		}
		if a.UpdatedAt != nil {
			t := a.UpdatedAt.AsTime()
			updatedAt = &t
		}

		accounts = append(accounts, models.Account{
			ID: a.Id,
			Email: a.Email,
			Name: a.Name,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		})
	}

	return accounts, nil
}
