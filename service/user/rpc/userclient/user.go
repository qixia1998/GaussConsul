// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: user.proto

package userclient

import (
	"context"

	"GaussConsul/service/user/rpc/types/user"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	DeleteUserRequest  = user.DeleteUserRequest
	DeleteUserResponse = user.DeleteUserResponse
	LoginRequest       = user.LoginRequest
	LoginResponse      = user.LoginResponse
	RegisterRequest    = user.RegisterRequest
	RegisterResponse   = user.RegisterResponse
	UpdateUserRequest  = user.UpdateUserRequest
	UpdateUserResponse = user.UpdateUserResponse
	UserInfoRequest    = user.UserInfoRequest
	UserInfoResponse   = user.UserInfoResponse

	User interface {
		Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
		Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
		UserInfo(ctx context.Context, in *UserInfoRequest, opts ...grpc.CallOption) (*UserInfoResponse, error)
		DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error)
		UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error)
	}

	defaultUser struct {
		cli zrpc.Client
	}
)

func NewUser(cli zrpc.Client) User {
	return &defaultUser{
		cli: cli,
	}
}

func (m *defaultUser) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.Login(ctx, in, opts...)
}

func (m *defaultUser) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.Register(ctx, in, opts...)
}

func (m *defaultUser) UserInfo(ctx context.Context, in *UserInfoRequest, opts ...grpc.CallOption) (*UserInfoResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.UserInfo(ctx, in, opts...)
}

func (m *defaultUser) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.DeleteUser(ctx, in, opts...)
}

func (m *defaultUser) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.UpdateUser(ctx, in, opts...)
}
