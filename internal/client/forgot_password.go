package client

import (
	"comet/utils"
	"context"
	"fmt"
	"google.golang.org/protobuf/types/known/emptypb"
	pb "protos/account"
)

// ForgotPassword forgot password request
func (cl *AccountClient) ForgotPassword(ctx context.Context) (*pb.ForgotPasswordResponse, error) {
	ctx, cancel, c, conn, err := cl.ConnectToAccountService(ctx)
	if err != nil {
		return nil, fmt.Errorf("connect to account service error: %w", err)
	}

	defer utils.HandleCloseConnection(conn)
	defer cancel()

	return c.ForgotPassword(ctx, &emptypb.Empty{})
}

// VerifyForgotPassword verify forgot password
func (cl *AccountClient) VerifyForgotPassword(ctx context.Context, forgotToken, code string) (*pb.VerifyForgotPasswordResponse, error) {
	ctx, cancel, c, conn, err := cl.ConnectToAccountService(ctx)
	if err != nil {
		return nil, fmt.Errorf("connect to account service error: %w", err)
	}

	defer utils.HandleCloseConnection(conn)
	defer cancel()

	return c.VerifyForgotPassword(ctx, &pb.VerifyForgotPasswordRequest{
		ForgotToken: forgotToken,
		Code:        code,
	})
}

// ResetPassword reset password
func (cl *AccountClient) ResetPassword(ctx context.Context, resetPasswordToken, newPassword string) (*pb.ResetPasswordResponse, error) {
	ctx, cancel, c, conn, err := cl.ConnectToAccountService(ctx)
	if err != nil {
		return nil, fmt.Errorf("connect to account service error: %w", err)
	}

	defer utils.HandleCloseConnection(conn)
	defer cancel()

	return c.ResetPassword(ctx, &pb.ResetPasswordRequest{
		ResetPasswordToken: resetPasswordToken,
		NewPassword:        newPassword,
	})
}
