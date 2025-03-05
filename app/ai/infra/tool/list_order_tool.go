package tool

import (
	"context"
	"fmt"
	"strconv"

	"github.com/SGNYYYY/gomall/app/ai/infra/rpc"
	"github.com/SGNYYYY/gomall/rpc_gen/kitex_gen/order"
	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/components/tool/utils"
)

// 参数结构体
type ListOrderReq struct {
	ID string `json:"id" jsonschema:"description=id of the todo"`
	// Content   *string `json:"content,omitempty" jsonschema:"description=content of the todo"`
	// StartedAt *int64  `json:"started_at,omitempty" jsonschema:"description=start time in unix timestamp"`
	// Deadline  *int64  `json:"deadline,omitempty" jsonschema:"description=deadline of the todo in unix timestamp"`
	// Done      *bool   `json:"done,omitempty" jsonschema:"description=done status"`
}

// 处理函数
func ListOrderFunc(ctx context.Context, req *ListOrderReq) (string, error) {
	userId, err := strconv.Atoi(req.ID)
	if err != nil {
		return `{"msg": "parse ID error"}`, err
	}
	resp, err := rpc.OrderClient.ListOrder(ctx, &order.ListOrderReq{
		UserId: uint32(userId),
	})
	if err != nil {
		return `{"msg": "server error"}`, err
	}
	var contentResp string
	contentResp = fmt.Sprintf("共查询到%d条订单信息\n", len(resp.Orders))
	for _, o := range resp.Orders {
		tmp := fmt.Sprintf("订单编号%s, 订单地址%s, 订单状态%s\n", o.OrderId, o.Address.Country+o.Address.State+o.Address.City+o.Address.StreetAddress, o.OrderState)
		contentResp += tmp
	}
	return contentResp, nil
}

func NewListOrderTool(_ context.Context) (tool.BaseTool, error) {
	// 使用 InferTool 创建工具
	listOrderTool, err := utils.InferTool(
		"query order",         // tool name
		"query user's orders", // tool description
		ListOrderFunc)
	return listOrderTool, err
}
