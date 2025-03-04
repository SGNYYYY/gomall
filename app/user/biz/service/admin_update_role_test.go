package service

import (
	"context"
	"testing"

	"github.com/SGNYYYY/gomall/app/user/biz/dal/mysql"
	user "github.com/SGNYYYY/gomall/rpc_gen/kitex_gen/user"
	"github.com/joho/godotenv"
)

func TestAdminUpdateRole_Run(t *testing.T) {
	_ = godotenv.Load("../../.env")
	mysql.Init()
	ctx := context.Background()
	s := NewAdminUpdateRoleService(ctx)
	// init req and assert value

	// 没有权限
	req := &user.AdminUpdateRoleReq{
		CurrentUserRole: "user",
		UserId:          2,
		Role:            "admin",
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// 用户不存在
	req = &user.AdminUpdateRoleReq{
		CurrentUserRole: "admin",
		UserId:          3,
		Role:            "admin",
	}
	resp, err = s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// 改为管理员
	req = &user.AdminUpdateRoleReq{
		CurrentUserRole: "admin",
		UserId:          2,
		Role:            "admin",
	}
	resp, err = s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// 改为用户
	req = &user.AdminUpdateRoleReq{
		CurrentUserRole: "admin",
		UserId:          2,
		Role:            "user",
	}
	resp, err = s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test
}
