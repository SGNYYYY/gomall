package service

import (
	"context"
	"errors"

	"github.com/SGNYYYY/gomall/app/user/biz/dal/mysql"
	"github.com/SGNYYYY/gomall/app/user/biz/model"
	user "github.com/SGNYYYY/gomall/rpc_gen/kitex_gen/user"
)

type AdminUpdateRoleService struct {
	ctx context.Context
} // NewAdminUpdateRoleService new AdminUpdateRoleService
func NewAdminUpdateRoleService(ctx context.Context) *AdminUpdateRoleService {
	return &AdminUpdateRoleService{ctx: ctx}
}

// Run create note info
func (s *AdminUpdateRoleService) Run(req *user.AdminUpdateRoleReq) (resp *user.AdminUpdateRoleResp, err error) {
	// Finish your business logic.
	if req.CurrentUserRole != string(model.RoleAdmin) {
		return nil, errors.New("没有权限")
	}
	var role model.UserRole
	if req.Role == string(model.RoleAdmin) {
		role = model.RoleAdmin
	} else {
		role = model.RoleUser
	}
	err = model.UpdateRole(mysql.DB, s.ctx, uint(req.UserId), role)
	if err != nil {
		return nil, err
	}
	return
}
