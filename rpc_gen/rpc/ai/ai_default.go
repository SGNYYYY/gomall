package ai

import (
	"context"
	ai "github.com/SGNYYYY/gomall/rpc_gen/kitex_gen/ai"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func Tool(ctx context.Context, req *ai.ToolReq, callOptions ...callopt.Option) (resp *ai.ToolResp, err error) {
	resp, err = defaultClient.Tool(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "Tool call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
