package user

import (
	"context"
	user "github.com/SGNYYYY/gomall/rpc_gen/kitex_gen/user"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func Register(ctx context.Context, req *user.RegisterReq, callOptions ...callopt.Option) (resp *user.RegisterResp, err error) {
	resp, err = defaultClient.Register(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "Register call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func Login(ctx context.Context, req *user.LoginReq, callOptions ...callopt.Option) (resp *user.LoginResp, err error) {
	resp, err = defaultClient.Login(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "Login call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func DeleteUser(ctx context.Context, req *user.DeleteUserReq, callOptions ...callopt.Option) (resp *user.DeleteUserResp, err error) {
	resp, err = defaultClient.DeleteUser(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "DeleteUser call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func UpdatePassword(ctx context.Context, req *user.UpdatePasswordReq, callOptions ...callopt.Option) (resp *user.UpdatePasswordResp, err error) {
	resp, err = defaultClient.UpdatePassword(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "UpdatePassword call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func AdminUpdatePassword(ctx context.Context, req *user.AdminUpdatePasswordReq, callOptions ...callopt.Option) (resp *user.AdminUpdatePasswordResp, err error) {
	resp, err = defaultClient.AdminUpdatePassword(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "AdminUpdatePassword call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func AdminUpdateRole(ctx context.Context, req *user.AdminUpdateRoleReq, callOptions ...callopt.Option) (resp *user.AdminUpdateRoleResp, err error) {
	resp, err = defaultClient.AdminUpdateRole(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "AdminUpdateRole call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func Activate(ctx context.Context, req *user.ActivateUserReq, callOptions ...callopt.Option) (resp *user.ActivateUserResp, err error) {
	resp, err = defaultClient.Activate(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "Activate call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func Deactivate(ctx context.Context, req *user.DeactivateUserReq, callOptions ...callopt.Option) (resp *user.DeactivateUserResp, err error) {
	resp, err = defaultClient.Deactivate(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "Deactivate call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func GetUserInfo(ctx context.Context, req *user.GetUserInfoReq, callOptions ...callopt.Option) (resp *user.GetUserInfoResp, err error) {
	resp, err = defaultClient.GetUserInfo(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GetUserInfo call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func GetUsers(ctx context.Context, req *user.GetUsersReq, callOptions ...callopt.Option) (resp *user.GetUsersResp, err error) {
	resp, err = defaultClient.GetUsers(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GetUsers call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
