package service

import (
	"context"
	"errors"

	"github.com/SGNYYYY/gomall/app/user/biz/dal/mysql"
	"github.com/SGNYYYY/gomall/app/user/biz/model"
	user "github.com/SGNYYYY/gomall/rpc_gen/kitex_gen/user"
)

type DeleteUserService struct {
	ctx context.Context
} // NewDeleteUserService new DeleteUserService
func NewDeleteUserService(ctx context.Context) *DeleteUserService {
	return &DeleteUserService{ctx: ctx}
}

// Run create note info
func (s *DeleteUserService) Run(req *user.DeleteUserReq) (resp *user.DeleteUserResp, err error) {
	if req.CurrentUserRole != string(model.RoleAdmin) {
		return nil, errors.New("没有权限")
	}
	err = model.DeleteUser(mysql.DB, s.ctx, uint(req.UserId))
	return
}
