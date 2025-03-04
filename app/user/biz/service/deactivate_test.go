package service

import (
	"context"
	"testing"

	"github.com/SGNYYYY/gomall/app/user/biz/dal/mysql"
	user "github.com/SGNYYYY/gomall/rpc_gen/kitex_gen/user"
	"github.com/joho/godotenv"
)

func TestDeactivate_Run(t *testing.T) {
	_ = godotenv.Load("../../.env")
	mysql.Init()
	ctx := context.Background()
	s := NewDeactivateService(ctx)
	// init req and assert value

	// 没有权限
	req := &user.DeactivateUserReq{
		CurrentUserRole: "user",
		UserId:          2,
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// 禁用用户
	req = &user.DeactivateUserReq{
		CurrentUserRole: "admin",
		UserId:          2,
	}
	resp, err = s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)
}
