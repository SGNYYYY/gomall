package service

import (
	"context"
	"errors"

	"github.com/SGNYYYY/gomall/app/user/biz/dal/mysql"
	"github.com/SGNYYYY/gomall/app/user/biz/model"
	user "github.com/SGNYYYY/gomall/rpc_gen/kitex_gen/user"
)

type ActivateService struct {
	ctx context.Context
} // NewActivateService new ActivateService
func NewActivateService(ctx context.Context) *ActivateService {
	return &ActivateService{ctx: ctx}
}

// Run create note info
func (s *ActivateService) Run(req *user.ActivateUserReq) (resp *user.ActivateUserResp, err error) {
	if req.CurrentUserRole != string(model.RoleAdmin) {
		return nil, errors.New("没有权限")
	}
	err = model.ActivateUser(mysql.DB, s.ctx, uint(req.UserId))
	if err != nil {
		return nil, err
	}
	return
}
