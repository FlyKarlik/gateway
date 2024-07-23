package client

import (
	"comet/utils"
	"context"
	"fmt"
	"google.golang.org/protobuf/types/known/emptypb"
	pb "protos/account"
)

// CheckUsername check username existence
func (cl *AccountClient) CheckUsername(ctx context.Context, username string) (*emptypb.Empty, error) {
	ctx, cancel, c, conn, err := cl.ConnectToAccountService(ctx)
	if err != nil {
		return nil, fmt.Errorf("connect to account service error: %w", err)
	}

	defer utils.HandleCloseConnection(conn)
	defer cancel()

	return c.CheckUsername(ctx, &pb.CheckUsernameRequest{
		Username: username,
	})
}

// GetAccountsInfo get account info
func (cl *AccountClient) GetAccountsInfo(ctx context.Context) (*pb.GetAccountsInfoResponse, error) {
	ctx, cancel, c, conn, err := cl.ConnectToAccountService(ctx)
	if err != nil {
		return nil, fmt.Errorf("connect to account service error: %w", err)
	}

	defer utils.HandleCloseConnection(conn)
	defer cancel()

	return c.GetAccountsInfo(ctx, &emptypb.Empty{})
}

// GetAccountInfo get account info
func (cl *AccountClient) GetAccountInfo(ctx context.Context) (*pb.GetAccountInfoResponse, error) {
	ctx, cancel, c, conn, err := cl.ConnectToAccountService(ctx)
	if err != nil {
		return nil, fmt.Errorf("connect to account service error: %w", err)
	}

	defer utils.HandleCloseConnection(conn)
	defer cancel()

	return c.GetAccountInfo(ctx, &emptypb.Empty{})
}
