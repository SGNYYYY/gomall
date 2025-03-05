package service

import (
	"context"

	"github.com/SGNYYYY/gomall/app/ai/infra/agent"
	ai "github.com/SGNYYYY/gomall/rpc_gen/kitex_gen/ai"
)

type ToolService struct {
	ctx context.Context
} // NewToolService new ToolService
func NewToolService(ctx context.Context) *ToolService {
	return &ToolService{ctx: ctx}
}

// Run create note info
func (s *ToolService) Run(req *ai.ToolReq) (resp *ai.ToolResp, err error) {
	// Finish your business logic.
	agent.Run(s.ctx, req.Content, req.UserId)
	return
}
