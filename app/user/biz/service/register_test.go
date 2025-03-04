package service

import (
	"context"
	"testing"

	"github.com/SGNYYYY/gomall/app/user/biz/dal/mysql"
	user "github.com/SGNYYYY/gomall/rpc_gen/kitex_gen/user"
	"github.com/joho/godotenv"
)

func TestRegister_Run(t *testing.T) {
	_ = godotenv.Load("../../.env")
	mysql.Init()

	ctx := context.Background()
	s := NewRegisterService(ctx)
	// init req and assert value

	req := &user.RegisterReq{
		Email:           "sgny@demo.com",
		Password:        "123456",
		ConfirmPassword: "123456",
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	req = &user.RegisterReq{
		Email:           "sgny@demo.com",
		Password:        "123456",
		ConfirmPassword: "123321",
	}
	resp, err = s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	req = &user.RegisterReq{
		Email:           "sgny@demo.com",
		Password:        "123456",
		ConfirmPassword: "123456",
	}
	resp, err = s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	req = &user.RegisterReq{
		Email:           "sgny2@demo.com",
		Password:        "123456",
		ConfirmPassword: "123456",
	}
	resp, err = s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)
}
