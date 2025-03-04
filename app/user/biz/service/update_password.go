package service

import (
	"context"
	"errors"

	"github.com/SGNYYYY/gomall/app/user/biz/dal/mysql"
	"github.com/SGNYYYY/gomall/app/user/biz/model"
	user "github.com/SGNYYYY/gomall/rpc_gen/kitex_gen/user"
	"golang.org/x/crypto/bcrypt"
)

type UpdatePasswordService struct {
	ctx context.Context
} // NewUpdatePasswordService new UpdatePasswordService
func NewUpdatePasswordService(ctx context.Context) *UpdatePasswordService {
	return &UpdatePasswordService{ctx: ctx}
}

// Run create note info
func (s *UpdatePasswordService) Run(req *user.UpdatePasswordReq) (resp *user.UpdatePasswordResp, err error) {
	if req.OldPassword == "" || req.Password == "" {
		return nil, errors.New("密码为空")
	}
	row, err := model.GetById(mysql.DB, s.ctx, uint(req.UserId))
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(row.PasswordHashed), []byte(req.OldPassword))
	if err != nil {
		return nil, err
	}
	newPasswordHashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	err = model.UpdatePassword(mysql.DB, s.ctx, uint(req.UserId), string(newPasswordHashed))
	if err != nil {
		return nil, err
	}
	return
}
