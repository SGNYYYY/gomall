package service

import (
	"context"
	"testing"

	"github.com/SGNYYYY/gomall/app/user/biz/dal/mysql"
	user "github.com/SGNYYYY/gomall/rpc_gen/kitex_gen/user"
	"github.com/joho/godotenv"
)

func TestAdminUpdatePassword_Run(t *testing.T) {
	_ = godotenv.Load("../../.env")
	mysql.Init()
	ctx := context.Background()
	s := NewAdminUpdatePasswordService(ctx)
	// init req and assert value

	// 没有权限
	req := &user.AdminUpdatePasswordReq{
		CurrentUserRole: "user",
		UserId:          2,
		Password:        "654321",
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// 更新密码
	req = &user.AdminUpdatePasswordReq{
		CurrentUserRole: "admin",
		UserId:          2,
		Password:        "654321",
	}
	resp, err = s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test
}
