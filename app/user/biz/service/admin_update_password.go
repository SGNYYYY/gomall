package service

import (
	"context"
	"errors"

	"github.com/SGNYYYY/gomall/app/user/biz/dal/mysql"
	"github.com/SGNYYYY/gomall/app/user/biz/model"
	user "github.com/SGNYYYY/gomall/rpc_gen/kitex_gen/user"
	"golang.org/x/crypto/bcrypt"
)

type AdminUpdatePasswordService struct {
	ctx context.Context
} // NewAdminUpdatePasswordService new AdminUpdatePasswordService
func NewAdminUpdatePasswordService(ctx context.Context) *AdminUpdatePasswordService {
	return &AdminUpdatePasswordService{ctx: ctx}
}

// Run create note info
func (s *AdminUpdatePasswordService) Run(req *user.AdminUpdatePasswordReq) (resp *user.AdminUpdatePasswordResp, err error) {
	if req.CurrentUserRole != string(model.RoleAdmin) {
		return nil, errors.New("没有权限")
	}
	passwordHashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	err = model.UpdatePassword(mysql.DB, s.ctx, uint(req.UserId), string(passwordHashed))
	if err != nil {
		return nil, err
	}
	return
}
