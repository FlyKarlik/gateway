package client

import (
	"comet/utils"
	"context"
	"fmt"
	pb "protos/account"
)

// Authorize with email
func (cl *AccountClient) Authorize(ctx context.Context, email string) (*pb.AuthorizeResponse, error) {
	ctx, cancel, c, conn, err := cl.ConnectToAccountService(ctx)
	if err != nil {
		return nil, fmt.Errorf("connect to account service error: %w", err)
	}

	defer utils.HandleCloseConnection(conn)
	defer cancel()

	return c.Authorize(ctx, &pb.AuthorizeRequest{
		Email: email,
	})
}

// VerifyAuthorize verify authorize
func (cl *AccountClient) VerifyAuthorize(ctx context.Context, authorizeToken, code string) (*pb.VerifyAuthorizeResponse, error) {
	ctx, cancel, c, conn, err := cl.ConnectToAccountService(ctx)
	if err != nil {
		return nil, fmt.Errorf("connect to account service error: %w", err)
	}

	defer utils.HandleCloseConnection(conn)
	defer cancel()

	return c.VerifyAuthorize(ctx, &pb.VerifyAuthorizeRequest{
		AuthorizeToken: authorizeToken,
		Code:           code,
	})
}

func (cl *AccountClient) LoginUser(ctx context.Context, email, password string) (*pb.LoginUserResponse, error) {
	ctx, cancel, c, conn, err := cl.ConnectToAccountService(ctx)
	if err != nil {
		return nil, fmt.Errorf("connect to account service error: %w", err)
	}

	defer utils.HandleCloseConnection(conn)
	defer cancel()

	return c.LoginUser(ctx, &pb.LoginUserRequest{
		Email:    email,
		Password: password})
}
