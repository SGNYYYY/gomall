package ai

import (
	"context"
	ai "github.com/SGNYYYY/gomall/rpc_gen/kitex_gen/ai"

	"github.com/SGNYYYY/gomall/rpc_gen/kitex_gen/ai/aiservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

type RPCClient interface {
	KitexClient() aiservice.Client
	Service() string
	Tool(ctx context.Context, Req *ai.ToolReq, callOptions ...callopt.Option) (r *ai.ToolResp, err error)
}

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := aiservice.NewClient(dstService, opts...)
	if err != nil {
		return nil, err
	}
	cli := &clientImpl{
		service:     dstService,
		kitexClient: kitexClient,
	}

	return cli, nil
}

type clientImpl struct {
	service     string
	kitexClient aiservice.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() aiservice.Client {
	return c.kitexClient
}

func (c *clientImpl) Tool(ctx context.Context, Req *ai.ToolReq, callOptions ...callopt.Option) (r *ai.ToolResp, err error) {
	return c.kitexClient.Tool(ctx, Req, callOptions...)
}
