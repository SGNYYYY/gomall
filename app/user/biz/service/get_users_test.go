package service

import (
	"context"
	"testing"

	"github.com/SGNYYYY/gomall/app/user/biz/dal/mysql"
	user "github.com/SGNYYYY/gomall/rpc_gen/kitex_gen/user"
	"github.com/joho/godotenv"
)

func TestGetUsers_Run(t *testing.T) {
	_ = godotenv.Load("../../.env")
	mysql.Init()
	ctx := context.Background()
	s := NewGetUsersService(ctx)
	// init req and assert value

	// 获取所有用户
	req := &user.GetUsersReq{
		CurrentUserRole: "admin",
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test
}
