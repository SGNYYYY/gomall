package service

import (
	"context"
	"testing"

	"github.com/SGNYYYY/gomall/app/user/biz/dal/mysql"
	user "github.com/SGNYYYY/gomall/rpc_gen/kitex_gen/user"
	"github.com/joho/godotenv"
)

func TestUpdatePassword_Run(t *testing.T) {
	_ = godotenv.Load("../../.env")
	mysql.Init()
	ctx := context.Background()
	s := NewUpdatePasswordService(ctx)
	// init req and assert value

	// 旧密码不正确
	req := &user.UpdatePasswordReq{
		UserId:      2,
		OldPassword: "123123",
		Password:    "321321",
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// 更新成功
	req = &user.UpdatePasswordReq{
		UserId:      2,
		OldPassword: "654321",
		Password:    "321321",
	}
	resp, err = s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test
}
