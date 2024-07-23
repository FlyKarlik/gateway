package client

import (
	"comet/utils"
	"context"
	"fmt"
	pb "protos/account"
)

// RefreshToken refresh token
func (cl *AccountClient) RefreshToken(ctx context.Context, refreshToken string) (*pb.RefreshTokenResponse, error) {
	ctx, cancel, c, conn, err := cl.ConnectToAccountService(ctx)
	if err != nil {
		return nil, fmt.Errorf("connect to account service error: %w", err)
	}

	defer utils.HandleCloseConnection(conn)
	defer cancel()

	return c.RefreshToken(ctx, &pb.RefreshTokenRequest{
		RefreshToken: refreshToken,
	})
}
