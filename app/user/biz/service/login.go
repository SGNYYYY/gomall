package service

import (
	"context"
	"errors"

	"github.com/SGNYYYY/gomall/app/user/biz/dal/mysql"
	"github.com/SGNYYYY/gomall/app/user/biz/model"
	user "github.com/SGNYYYY/gomall/rpc_gen/kitex_gen/user"
	"golang.org/x/crypto/bcrypt"
)

type LoginService struct {
	ctx context.Context
} // NewLoginService new LoginService
func NewLoginService(ctx context.Context) *LoginService {
	return &LoginService{ctx: ctx}
}

// Run create note info
func (s *LoginService) Run(req *user.LoginReq) (resp *user.LoginResp, err error) {
	// Finish your business logic.
	if req.Email == "" || req.Password == "" {
		return nil, errors.New("用户名或密码为空")
	}
	row, err := model.GetByEmail(mysql.DB, s.ctx, req.Email)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, errors.New("账号不存在")
	}
	err = bcrypt.CompareHashAndPassword([]byte(row.PasswordHashed), []byte(req.Password))
	if err != nil {
		return nil, err
	}
	if row.Status == model.Disabled {
		return nil, errors.New("账号已禁用")
	}
	resp = &user.LoginResp{
		UserId: int32(row.ID),
		Role:   string(row.Role),
	}
	return resp, nil
}
