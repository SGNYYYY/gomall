package main

import (
	"context"

	"github.com/SGNYYYY/gomall/app/user/biz/service"
	user "github.com/SGNYYYY/gomall/rpc_gen/kitex_gen/user"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	resp, err = service.NewRegisterService(ctx).Run(req)

	return resp, err
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *user.LoginReq) (resp *user.LoginResp, err error) {
	resp, err = service.NewLoginService(ctx).Run(req)

	return resp, err
}

// DeleteUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) DeleteUser(ctx context.Context, req *user.DeleteUserReq) (resp *user.DeleteUserResp, err error) {
	resp, err = service.NewDeleteUserService(ctx).Run(req)

	return resp, err
}

// GetUserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUserInfo(ctx context.Context, req *user.GetUserInfoReq) (resp *user.GetUserInfoResp, err error) {
	resp, err = service.NewGetUserInfoService(ctx).Run(req)

	return resp, err
}

// GetUsers implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUsers(ctx context.Context, req *user.GetUsersReq) (resp *user.GetUsersResp, err error) {
	resp, err = service.NewGetUsersService(ctx).Run(req)

	return resp, err
}

// UpdatePassword implements the UserServiceImpl interface.
func (s *UserServiceImpl) UpdatePassword(ctx context.Context, req *user.UpdatePasswordReq) (resp *user.UpdatePasswordResp, err error) {
	resp, err = service.NewUpdatePasswordService(ctx).Run(req)

	return resp, err
}

// AdminUpdatePassword implements the UserServiceImpl interface.
func (s *UserServiceImpl) AdminUpdatePassword(ctx context.Context, req *user.AdminUpdatePasswordReq) (resp *user.AdminUpdatePasswordResp, err error) {
	resp, err = service.NewAdminUpdatePasswordService(ctx).Run(req)

	return resp, err
}

// AdminUpdateRole implements the UserServiceImpl interface.
func (s *UserServiceImpl) AdminUpdateRole(ctx context.Context, req *user.AdminUpdateRoleReq) (resp *user.AdminUpdateRoleResp, err error) {
	resp, err = service.NewAdminUpdateRoleService(ctx).Run(req)

	return resp, err
}

// Activate implements the UserServiceImpl interface.
func (s *UserServiceImpl) Activate(ctx context.Context, req *user.ActivateUserReq) (resp *user.ActivateUserResp, err error) {
	resp, err = service.NewActivateService(ctx).Run(req)

	return resp, err
}

// Deactivate implements the UserServiceImpl interface.
func (s *UserServiceImpl) Deactivate(ctx context.Context, req *user.DeactivateUserReq) (resp *user.DeactivateUserResp, err error) {
	resp, err = service.NewDeactivateService(ctx).Run(req)

	return resp, err
}
