package interfaces

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	pb "protos/account"
)

type AccountClient interface {
	Authorize(ctx context.Context, email string) (*pb.AuthorizeResponse, error)
	VerifyAuthorize(ctx context.Context, authorizeToken, code string) (*pb.VerifyAuthorizeResponse, error)
	RegisterPassword(ctx context.Context, password string) (*emptypb.Empty, error)
	RegisterUsername(ctx context.Context, username string) (*pb.RegisterUsernameResponse, error)
	CheckUsername(ctx context.Context, username string) (*emptypb.Empty, error)
	ForgotPassword(ctx context.Context) (*pb.ForgotPasswordResponse, error)
	VerifyForgotPassword(ctx context.Context, forgotToken, code string) (*pb.VerifyForgotPasswordResponse, error)
	ResetPassword(ctx context.Context, resetPasswordToken, newPassword string) (*pb.ResetPasswordResponse, error)
	RefreshToken(ctx context.Context, refreshToken string) (*pb.RefreshTokenResponse, error)
	GetAccountInfo(ctx context.Context) (*pb.GetAccountInfoResponse, error)
	GetAccountsInfo(ctx context.Context) (*pb.GetAccountsInfoResponse, error)
	RegisterUser(ctx context.Context, email, firstName, secondName, password string, departmentID, roleID uint32) (*pb.RegisterUserResponse, error)
	RemoveUser(ctx context.Context, id string) (*emptypb.Empty, error)
	AddUserRole(ctx context.Context, name string) (*pb.AddUserRoleResponse, error)
	RemoveUserRole(ctx context.Context, id uint32) (*emptypb.Empty, error)
	GetUserRoles(ctx context.Context) (*pb.GetUserRolesResponse, error)
	GetUserRole(ctx context.Context, id uint32) (*pb.GetUserRoleResponse, error)
	AddUserDepartment(ctx context.Context, name string) (*pb.AddUserDepartmentResponse, error)
	RemoveUserDepartment(ctx context.Context, id uint32) (*emptypb.Empty, error)
	GetUserDepartments(ctx context.Context) (*pb.GetUserDepartmentsResponse, error)
	GetUserDepartment(ctx context.Context, id uint32) (*pb.GetUserDepartmentResponse, error)
	LoginUser(ctx context.Context, email, password string) (*pb.LoginUserResponse, error)
}
