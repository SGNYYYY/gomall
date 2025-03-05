package service

import (
	"context"
	"testing"

	"github.com/SGNYYYY/gomall/app/ai/infra/rpc"
	ai "github.com/SGNYYYY/gomall/rpc_gen/kitex_gen/ai"
	"github.com/joho/godotenv"
)

func TestTool_Run(t *testing.T) {
	_ = godotenv.Load("../../.env")
	rpc.InitClient()
	ctx := context.Background()
	s := NewToolService(ctx)
	// init req and assert value
	var req *ai.ToolReq
	var resp *ai.ToolResp
	var err error
	req = &ai.ToolReq{
		UserId:  10,
		Content: "查找我的订单",
	}
	resp, err = s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	req = &ai.ToolReq{
		UserId:  10,
		Content: "下单购物车里的商品,我的地址在中国广东省广州市小谷围街道,邮政编码100000",
	}
	resp, err = s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test
}
