package internal

import (
	"aurora-graph/account/proto/pb"
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type grpcServer struct {
	pb.UnimplementedAccountServiceServer
	service Service
}

func ListenGRPC(service Service, port int) error {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}

	server := grpc.NewServer()

	pb.RegisterAccountServiceServer(server, &grpcServer{
		UnimplementedAccountServiceServer: pb.UnimplementedAccountServiceServer{},
		service: service,
	})

	return server.Serve(listen)
}

func (server *grpcServer) Register(ctx context.Context, request *pb.RegisterRequest) (*wrapperspb.StringValue, error) {
	token, err := server.service.Register(ctx, request.Name, request.Email, request.Password)
	if err != nil {
		return nil, err
	}

	return &wrapperspb.StringValue{
		Value: token,
	}, nil
}

func (server *grpcServer) Login(ctx context.Context, request *pb.LoginRequest) (*wrapperspb.StringValue, error) {
	token, err := server.service.Login(ctx, request.Email, request.Password)
	if err != nil {
		return nil, err
	}

	return &wrapperspb.StringValue{
		Value: token,
	}, nil
}

func (server *grpcServer) GetAccount(ctx context.Context, request *pb.GetAccountRequest) (*pb.AccountResponse, error) {
	account, err := server.service.GetAccount(ctx, request.Id)
	if err != nil {
		return nil, err
	}

	return &pb.AccountResponse{
		Account: &pb.Account{
			Id: account.ID,
			Name: account.Name,
		},
	}, nil
}

func (server *grpcServer) GetAccounts(ctx context.Context, request *pb.GetAccountsRequest) (*pb.GetAccountsResponse, error) {
	result, err := server.service.GetAccounts(ctx, request.Skip, request.Take)
	if err != nil {
		return nil, err
	}

	var accounts []*pb.Account

	for _, p := range result {
		accounts = append(accounts, &pb.Account{
			Id: p.ID,
			Name: p.Name,
		})
	}

	return &pb.GetAccountsResponse{Accounts: accounts}, nil
} 