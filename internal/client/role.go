package client

import (
	"comet/utils"
	"context"
	"fmt"
	"google.golang.org/protobuf/types/known/emptypb"
	pb "protos/account"
)

// AddUserRole add new department
func (cl *AccountClient) AddUserRole(ctx context.Context, name string) (*pb.AddUserRoleResponse, error) {
	ctx, cancel, c, conn, err := cl.ConnectToAccountService(ctx)
	if err != nil {
		return nil, fmt.Errorf("connect to account service error: %w", err)
	}

	defer utils.HandleCloseConnection(conn)
	defer cancel()

	return c.AddUserRole(ctx, &pb.AddUserRoleRequest{
		Name: name,
	})
}

// RemoveUserRole remove department by id
func (cl *AccountClient) RemoveUserRole(ctx context.Context, id uint32) (*emptypb.Empty, error) {
	ctx, cancel, c, conn, err := cl.ConnectToAccountService(ctx)
	if err != nil {
		return nil, fmt.Errorf("connect to account service error: %w", err)
	}

	defer utils.HandleCloseConnection(conn)
	defer cancel()

	return c.RemoveUserRole(ctx, &pb.RemoveUserRoleRequest{Id: id})
}

// GetUserRoles get all departments
func (cl *AccountClient) GetUserRoles(ctx context.Context) (*pb.GetUserRolesResponse, error) {
	ctx, cancel, c, conn, err := cl.ConnectToAccountService(ctx)
	if err != nil {
		return nil, fmt.Errorf("connect to account service error: %w", err)
	}

	defer utils.HandleCloseConnection(conn)
	defer cancel()

	return c.GetUserRoles(ctx, &emptypb.Empty{})
}

// GetUserRole get department by ID
func (cl *AccountClient) GetUserRole(ctx context.Context, id uint32) (*pb.GetUserRoleResponse, error) {
	ctx, cancel, c, conn, err := cl.ConnectToAccountService(ctx)
	if err != nil {
		return nil, fmt.Errorf("connect to account service error: %w", err)
	}

	defer utils.HandleCloseConnection(conn)
	defer cancel()

	return c.GetUserRole(ctx, &pb.GetUserRoleRequest{Id: id})
}
