package client

import (
	"comet/utils"
	"context"
	"fmt"
	"google.golang.org/protobuf/types/known/emptypb"
	pb "protos/account"
)

// AddUserDepartment add new department
func (cl *AccountClient) AddUserDepartment(ctx context.Context, name string) (*pb.AddUserDepartmentResponse, error) {
	ctx, cancel, c, conn, err := cl.ConnectToAccountService(ctx)
	if err != nil {
		return nil, fmt.Errorf("connect to account service error: %w", err)
	}

	defer utils.HandleCloseConnection(conn)
	defer cancel()

	return c.AddUserDepartment(ctx, &pb.AddUserDepartmentRequest{
		Name: name,
	})
}

// RemoveUserDepartment remove department by id
func (cl *AccountClient) RemoveUserDepartment(ctx context.Context, id uint32) (*emptypb.Empty, error) {
	ctx, cancel, c, conn, err := cl.ConnectToAccountService(ctx)
	if err != nil {
		return nil, fmt.Errorf("connect to account service error: %w", err)
	}

	defer utils.HandleCloseConnection(conn)
	defer cancel()

	return c.RemoveUserDepartment(ctx, &pb.RemoveUserDepartmentRequest{Id: id})
}

// GetUserDepartments get all departments
func (cl *AccountClient) GetUserDepartments(ctx context.Context) (*pb.GetUserDepartmentsResponse, error) {
	ctx, cancel, c, conn, err := cl.ConnectToAccountService(ctx)
	if err != nil {
		return nil, fmt.Errorf("connect to account service error: %w", err)
	}

	defer utils.HandleCloseConnection(conn)
	defer cancel()

	return c.GetUserDepartments(ctx, &emptypb.Empty{})
}

// GetUserDepartment get department by ID
func (cl *AccountClient) GetUserDepartment(ctx context.Context, id uint32) (*pb.GetUserDepartmentResponse, error) {
	ctx, cancel, c, conn, err := cl.ConnectToAccountService(ctx)
	if err != nil {
		return nil, fmt.Errorf("connect to account service error: %w", err)
	}

	defer utils.HandleCloseConnection(conn)
	defer cancel()

	return c.GetUserDepartment(ctx, &pb.GetUserDepartmentRequest{Id: id})
}
