package service

import (
	"context"
	"errors"

	"github.com/SGNYYYY/gomall/app/user/biz/dal/mysql"
	"github.com/SGNYYYY/gomall/app/user/biz/model"
	user "github.com/SGNYYYY/gomall/rpc_gen/kitex_gen/user"
)

type GetUsersService struct {
	ctx context.Context
} // NewGetUsersService new GetUsersService
func NewGetUsersService(ctx context.Context) *GetUsersService {
	return &GetUsersService{ctx: ctx}
}

// Run create note info
func (s *GetUsersService) Run(req *user.GetUsersReq) (resp *user.GetUsersResp, err error) {
	if req.CurrentUserRole != string(model.RoleAdmin) {
		return nil, errors.New("没有权限")
	}
	u, err := model.GetUsers(mysql.DB, s.ctx)
	if err != nil {
		return nil, err
	}
	var users []*user.User
	for _, ui := range u {
		users = append(users, &user.User{
			Email:     ui.Email,
			Role:      string(ui.Role),
			CreatedAt: int32(ui.CreatedAt.Unix()),
			UpdatedAt: int32(ui.UpdatedAt.Unix()),
			Status:    string(ui.Status),
		})
	}
	return &user.GetUsersResp{
		Users: users,
	}, nil
}
