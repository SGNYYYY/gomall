package tool

import (
	"context"
	"fmt"
	"strconv"

	"github.com/SGNYYYY/gomall/app/ai/infra/rpc"
	"github.com/SGNYYYY/gomall/rpc_gen/kitex_gen/checkout"
	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/components/tool/utils"
)

// 参数结构体
type PlaceOrderReq struct {
	ID            string `json:"id" jsonschema:"description=id of the todo"`
	StreetAddress string `json:"street_address,omitempty" jsonschema:"description=street address of user"`
	City          string `json:"city,omitempty" jsonschema:"description=city of user's address"`
	State         string `json:"state,omitempty" jsonschema:"description=state of user's address"`
	Country       string `json:"country,omitempty" jsonschema:"description=country of user's address"`
	ZipCode       string `json:"zip_code,omitempty" jsonschema:"description=zipcode of user's address"`
}

// 处理函数
func PlaceOrderFunc(ctx context.Context, req *PlaceOrderReq) (string, error) {
	userId, err := strconv.Atoi(req.ID)
	if err != nil {
		return `{"msg": "parse ID error"}`, err
	}
	resp, err := rpc.CheckoutClient.Checkout(ctx, &checkout.CheckoutReq{
		UserId: uint32(userId),
		Address: &checkout.Address{
			StreetAddress: req.StreetAddress,
			City:          req.City,
			State:         req.State,
			Country:       req.Country,
			ZipCode:       req.ZipCode,
		},
	})
	if err != nil {
		return `{"msg": "server error"}`, err
	}
	fmt.Println(resp)
	return `{"msg": "下单成功!请尽快支付"}`, nil
}

func NewPlaceOrderTool(_ context.Context) (tool.BaseTool, error) {
	// 使用 InferTool 创建工具
	placeOrderTool, err := utils.InferTool(
		"place order",          // tool name
		"place order for user", // tool description
		PlaceOrderFunc)
	return placeOrderTool, err
}
