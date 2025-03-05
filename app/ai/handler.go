package main

import (
	"context"
	ai "github.com/SGNYYYY/gomall/rpc_gen/kitex_gen/ai"
	"github.com/SGNYYYY/gomall/app/ai/biz/service"
)

// AiServiceImpl implements the last service interface defined in the IDL.
type AiServiceImpl struct{}

// Tool implements the AiServiceImpl interface.
func (s *AiServiceImpl) Tool(ctx context.Context, req *ai.ToolReq) (resp *ai.ToolResp, err error) {
	resp, err = service.NewToolService(ctx).Run(req)

	return resp, err
}
