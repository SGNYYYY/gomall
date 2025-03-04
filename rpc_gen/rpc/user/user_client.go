package user

import (
	"context"
	user "github.com/SGNYYYY/gomall/rpc_gen/kitex_gen/user"

	"github.com/SGNYYYY/gomall/rpc_gen/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

type RPCClient interface {
	KitexClient() userservice.Client
	Service() string
	Register(ctx context.Context, Req *user.RegisterReq, callOptions ...callopt.Option) (r *user.RegisterResp, err error)
	Login(ctx context.Context, Req *user.LoginReq, callOptions ...callopt.Option) (r *user.LoginResp, err error)
	DeleteUser(ctx context.Context, Req *user.DeleteUserReq, callOptions ...callopt.Option) (r *user.DeleteUserResp, err error)
	UpdatePassword(ctx context.Context, Req *user.UpdatePasswordReq, callOptions ...callopt.Option) (r *user.UpdatePasswordResp, err error)
	AdminUpdatePassword(ctx context.Context, Req *user.AdminUpdatePasswordReq, callOptions ...callopt.Option) (r *user.AdminUpdatePasswordResp, err error)
	AdminUpdateRole(ctx context.Context, Req *user.AdminUpdateRoleReq, callOptions ...callopt.Option) (r *user.AdminUpdateRoleResp, err error)
	Activate(ctx context.Context, Req *user.ActivateUserReq, callOptions ...callopt.Option) (r *user.ActivateUserResp, err error)
	Deactivate(ctx context.Context, Req *user.DeactivateUserReq, callOptions ...callopt.Option) (r *user.DeactivateUserResp, err error)
	GetUserInfo(ctx context.Context, Req *user.GetUserInfoReq, callOptions ...callopt.Option) (r *user.GetUserInfoResp, err error)
	GetUsers(ctx context.Context, Req *user.GetUsersReq, callOptions ...callopt.Option) (r *user.GetUsersResp, err error)
}

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := userservice.NewClient(dstService, opts...)
	if err != nil {
		return nil, err
	}
	cli := &clientImpl{
		service:     dstService,
		kitexClient: kitexClient,
	}

	return cli, nil
}

type clientImpl struct {
	service     string
	kitexClient userservice.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() userservice.Client {
	return c.kitexClient
}

func (c *clientImpl) Register(ctx context.Context, Req *user.RegisterReq, callOptions ...callopt.Option) (r *user.RegisterResp, err error) {
	return c.kitexClient.Register(ctx, Req, callOptions...)
}

func (c *clientImpl) Login(ctx context.Context, Req *user.LoginReq, callOptions ...callopt.Option) (r *user.LoginResp, err error) {
	return c.kitexClient.Login(ctx, Req, callOptions...)
}

func (c *clientImpl) DeleteUser(ctx context.Context, Req *user.DeleteUserReq, callOptions ...callopt.Option) (r *user.DeleteUserResp, err error) {
	return c.kitexClient.DeleteUser(ctx, Req, callOptions...)
}

func (c *clientImpl) UpdatePassword(ctx context.Context, Req *user.UpdatePasswordReq, callOptions ...callopt.Option) (r *user.UpdatePasswordResp, err error) {
	return c.kitexClient.UpdatePassword(ctx, Req, callOptions...)
}

func (c *clientImpl) AdminUpdatePassword(ctx context.Context, Req *user.AdminUpdatePasswordReq, callOptions ...callopt.Option) (r *user.AdminUpdatePasswordResp, err error) {
	return c.kitexClient.AdminUpdatePassword(ctx, Req, callOptions...)
}

func (c *clientImpl) AdminUpdateRole(ctx context.Context, Req *user.AdminUpdateRoleReq, callOptions ...callopt.Option) (r *user.AdminUpdateRoleResp, err error) {
	return c.kitexClient.AdminUpdateRole(ctx, Req, callOptions...)
}

func (c *clientImpl) Activate(ctx context.Context, Req *user.ActivateUserReq, callOptions ...callopt.Option) (r *user.ActivateUserResp, err error) {
	return c.kitexClient.Activate(ctx, Req, callOptions...)
}

func (c *clientImpl) Deactivate(ctx context.Context, Req *user.DeactivateUserReq, callOptions ...callopt.Option) (r *user.DeactivateUserResp, err error) {
	return c.kitexClient.Deactivate(ctx, Req, callOptions...)
}

func (c *clientImpl) GetUserInfo(ctx context.Context, Req *user.GetUserInfoReq, callOptions ...callopt.Option) (r *user.GetUserInfoResp, err error) {
	return c.kitexClient.GetUserInfo(ctx, Req, callOptions...)
}

func (c *clientImpl) GetUsers(ctx context.Context, Req *user.GetUsersReq, callOptions ...callopt.Option) (r *user.GetUsersResp, err error) {
	return c.kitexClient.GetUsers(ctx, Req, callOptions...)
}
