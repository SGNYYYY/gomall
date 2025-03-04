package service

import (
	"context"

	"github.com/SGNYYYY/gomall/app/user/biz/dal/mysql"
	"github.com/SGNYYYY/gomall/app/user/biz/model"
	user "github.com/SGNYYYY/gomall/rpc_gen/kitex_gen/user"
)

type GetUserInfoService struct {
	ctx context.Context
} // NewGetUserInfoService new GetUserInfoService
func NewGetUserInfoService(ctx context.Context) *GetUserInfoService {
	return &GetUserInfoService{ctx: ctx}
}

// Run create note info
func (s *GetUserInfoService) Run(req *user.GetUserInfoReq) (resp *user.GetUserInfoResp, err error) {
	u, err := model.GetById(mysql.DB, s.ctx, uint(req.UserId))
	if err != nil {
		return nil, err
	}

	return &user.GetUserInfoResp{
		Email: u.Email,
	}, nil
}
