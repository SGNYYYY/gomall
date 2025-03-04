package service

import (
	"context"
	"testing"

	"github.com/SGNYYYY/gomall/app/user/biz/dal/mysql"
	user "github.com/SGNYYYY/gomall/rpc_gen/kitex_gen/user"
	"github.com/joho/godotenv"
)

func TestLogin_Run(t *testing.T) {
	_ = godotenv.Load("../../.env")
	mysql.Init()
	ctx := context.Background()
	s := NewLoginService(ctx)

	// 成功登录
	req := &user.LoginReq{
		Email:    "sgny@demo.com",
		Password: "123456",
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)
	// 密码错误
	req = &user.LoginReq{
		Email:    "sgny@demo.com",
		Password: "123321",
	}
	resp, err = s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)
	// 账号不存在
	req = &user.LoginReq{
		Email:    "sgny1@demo.com",
		Password: "123321",
	}
	resp, err = s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// 账号已禁用
	req = &user.LoginReq{
		Email:    "sgny2@demo.com",
		Password: "123456",
	}
	resp, err = s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)
}
