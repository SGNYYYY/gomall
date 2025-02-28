package service

import (
	"context"
	"fmt"
	"strconv"

	"github.com/SGNYYYY/gomall/app/checkout/infra/mq"
	"github.com/SGNYYYY/gomall/app/checkout/infra/rpc"
	rpccart "github.com/SGNYYYY/gomall/rpc_gen/kitex_gen/cart"
	checkout "github.com/SGNYYYY/gomall/rpc_gen/kitex_gen/checkout"
	rpcemail "github.com/SGNYYYY/gomall/rpc_gen/kitex_gen/email"
	rpcorder "github.com/SGNYYYY/gomall/rpc_gen/kitex_gen/order"
	rpcproduct "github.com/SGNYYYY/gomall/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/nats-io/nats.go"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"google.golang.org/protobuf/proto"
)

type CheckoutService struct {
	ctx context.Context
} // NewCheckoutService new CheckoutService
func NewCheckoutService(ctx context.Context) *CheckoutService {
	return &CheckoutService{ctx: ctx}
}

// Run create note info
func (s *CheckoutService) Run(req *checkout.CheckoutReq) (resp *checkout.CheckoutResp, err error) {
	// Finish your business logic.
	cartResult, err := rpc.CartClient.GetCart(s.ctx, &rpccart.GetCartReq{
		UserId: req.UserId,
	})
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(5005001, err.Error())
	}
	if cartResult == nil || cartResult.Cart.Items == nil {
		return nil, kerrors.NewGRPCBizStatusError(5004001, "cart is empty")
	}

	// 从购物车中获取订单item信息
	var oi []*rpcorder.OrderItem
	var total float32
	for _, cartItem := range cartResult.Cart.Items {
		// 避免在for循环内调用rpc服务, 应该在for循环外统一获取然后遍历赋值
		productResp, resultErr := rpc.ProductClient.GetProduct(s.ctx, &rpcproduct.GetProductReq{
			Id: cartItem.ProductId,
		})

		if resultErr != nil {
			return nil, resultErr
		}

		if productResp.Product == nil {
			continue
		}

		p := productResp.Product.Price

		cost := float32(cartItem.Quantity) * p
		oi = append(oi, &rpcorder.OrderItem{
			Cost: cost,
			Item: &rpccart.CartItem{ProductId: cartItem.ProductId, Quantity: cartItem.Quantity},
		})
		total += cost
	}

	// 下订单，下单的时候支付的状态为未支付
	var orderId string
	orderReq := &rpcorder.PlaceOrderReq{
		UserId:       req.UserId,
		UserCurrency: "USD",
		OrderItems:   oi,
		Email:        req.Email,
	}
	if req.Address != nil {
		addr := req.Address
		zipCodeInt, _ := strconv.Atoi(addr.ZipCode)
		orderReq.Address = &rpcorder.Address{
			StreetAddress: addr.StreetAddress,
			City:          addr.City,
			Country:       addr.Country,
			State:         addr.State,
			ZipCode:       int32(zipCodeInt),
		}
	}
	orderResult, err := rpc.OrderClient.PlaceOrder(s.ctx, orderReq)
	if err != nil {
		err = fmt.Errorf("PlaceOrder.err:%v", err)
		return
	}
	klog.Info("orderResult", orderResult)
	if orderResult != nil && orderResult.Order != nil {
		orderId = orderResult.Order.OrderId
	}

	// TODO，把支付放到支付里面
	// payReq := &payment.ChargeReq{
	// 	UserId:  req.UserId,
	// 	OrderId: orderId,
	// 	Amount:  total,
	// 	CreditCard: &payment.CreditCardInfo{
	// 		CreditCardNumber:          req.CreditCard.CreditCardNumber,
	// 		CreditCardCvv:             req.CreditCard.CreditCardCvv,
	// 		CreditCardExpirationMonth: req.CreditCard.CreditCardExpirationMonth,
	// 		CreditCardExpirationYear:  req.CreditCard.CreditCardExpirationYear,
	// 	},
	// }
	// 清空购物车
	_, err = rpc.CartClient.EmptyCart(s.ctx, &rpccart.EmptyCartReq{UserId: req.UserId})
	if err != nil {
		klog.Error(err.Error())
	}

	data, _ := proto.Marshal(&rpcemail.EmailReq{
		From:        "from@example.com",
		To:          req.Email,
		ContentType: "text/plain",
		Subject:     "You have just created an order in GOMALL",
		Content:     "You have just created an order in GOMALL",
	})
	msg := &nats.Msg{Subject: "email", Data: data, Header: make(nats.Header)}
	otel.GetTextMapPropagator().Inject(s.ctx, propagation.HeaderCarrier(msg.Header))

	_ = mq.Nc.PublishMsg(msg)

	// paymentResult, err := rpc.PaymentClient.Charge(s.ctx, payReq)

	// if err != nil {
	// 	return nil, err
	// }

	// klog.Info(paymentResult)
	resp = &checkout.CheckoutResp{
		OrderId: orderId,
		// TransactionId: paymentResult.TransactionId,
	}
	return resp, nil
}
