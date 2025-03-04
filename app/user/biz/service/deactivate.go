package service

import (
	"context"
	"errors"

	"github.com/SGNYYYY/gomall/app/user/biz/dal/mysql"
	"github.com/SGNYYYY/gomall/app/user/biz/model"
	user "github.com/SGNYYYY/gomall/rpc_gen/kitex_gen/user"
)

type DeactivateService struct {
	ctx context.Context
} // NewDeactivateService new DeactivateService
func NewDeactivateService(ctx context.Context) *DeactivateService {
	return &DeactivateService{ctx: ctx}
}

// Run create note info
func (s *DeactivateService) Run(req *user.DeactivateUserReq) (resp *user.DeactivateUserResp, err error) {
	if req.CurrentUserRole != string(model.RoleAdmin) {
		return nil, errors.New("没有权限")
	}
	err = model.DeactivateUser(mysql.DB, s.ctx, uint(req.UserId))
	if err != nil {
		return nil, err
	}
	return
}
