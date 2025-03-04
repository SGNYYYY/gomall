package service

import (
	"context"
	"testing"

	"github.com/SGNYYYY/gomall/app/user/biz/dal/mysql"
	user "github.com/SGNYYYY/gomall/rpc_gen/kitex_gen/user"
	"github.com/joho/godotenv"
)

func TestDeleteUser_Run(t *testing.T) {
	_ = godotenv.Load("../../.env")
	mysql.Init()
	ctx := context.Background()
	s := NewDeleteUserService(ctx)
	// init req and assert value

	// 没有权限
	req := &user.DeleteUserReq{
		CurrentUserRole: "user",
		UserId:          8,
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// 用户不存在
	req = &user.DeleteUserReq{
		CurrentUserRole: "admin",
		UserId:          9,
	}
	resp, err = s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)
	// 成功删除
	req = &user.DeleteUserReq{
		CurrentUserRole: "admin",
		UserId:          8,
	}
	resp, err = s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test
}
