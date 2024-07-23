package client

import (
	"comet/utils"
	"context"
	"fmt"
	"google.golang.org/protobuf/types/known/emptypb"
	pb "protos/account"
)

// RegisterPassword register password
func (cl *AccountClient) RegisterPassword(ctx context.Context, password string) (*emptypb.Empty, error) {
	ctx, cancel, c, conn, err := cl.ConnectToAccountService(ctx)
	if err != nil {
		return nil, fmt.Errorf("connect to account service error: %w", err)
	}
	defer utils.HandleCloseConnection(conn)
	defer cancel()

	return c.RegisterPassword(ctx, &pb.RegisterPasswordRequest{
		Password: password,
	})
}

// RegisterUsername register username
func (cl *AccountClient) RegisterUsername(ctx context.Context, username string) (*pb.RegisterUsernameResponse, error) {
	ctx, cancel, c, conn, err := cl.ConnectToAccountService(ctx)
	if err != nil {
		return nil, fmt.Errorf("connect to account service error: %w", err)
	}

	defer utils.HandleCloseConnection(conn)
	defer cancel()

	return c.RegisterUsername(ctx, &pb.RegisterUsernameRequest{
		Username: username,
	})
}

// RegisterUser register user
func (cl *AccountClient) RegisterUser(ctx context.Context, email, firstName, secondName, password string, departmentID, roleID uint32) (*pb.RegisterUserResponse, error) {
	ctx, cancel, c, conn, err := cl.ConnectToAccountService(ctx)
	if err != nil {
		return nil, fmt.Errorf("connect to account service error: %w", err)
	}

	defer utils.HandleCloseConnection(conn)
	defer cancel()

	return c.RegisterUser(ctx, &pb.RegisterUserRequest{
		Email:        email,
		FirstName:    firstName,
		SecondName:   secondName,
		Password:     password,
		DepartmentId: departmentID,
		RoleId:       roleID,
	})
}

func (cl *AccountClient) RemoveUser(ctx context.Context, id string) (*emptypb.Empty, error) {
	return nil, nil
}
