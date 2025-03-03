package rpc

import (
	"context"
	"sync"

	"github.com/SGNYYYY/gomall/app/frontend/conf"
	frontendUtils "github.com/SGNYYYY/gomall/app/frontend/utils"
	"github.com/SGNYYYY/gomall/common/clientsuite"
	"github.com/SGNYYYY/gomall/rpc_gen/kitex_gen/auth/authservice"
	"github.com/SGNYYYY/gomall/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/SGNYYYY/gomall/rpc_gen/kitex_gen/checkout/checkoutservice"
	"github.com/SGNYYYY/gomall/rpc_gen/kitex_gen/order/orderservice"
	"github.com/SGNYYYY/gomall/rpc_gen/kitex_gen/payment/paymentservice"
	"github.com/SGNYYYY/gomall/rpc_gen/kitex_gen/product"
	"github.com/SGNYYYY/gomall/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/SGNYYYY/gomall/rpc_gen/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/circuitbreak"
	"github.com/cloudwego/kitex/pkg/fallback"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
)

var (
	UserClient     userservice.Client
	ProductClient  productcatalogservice.Client
	CartClient     cartservice.Client
	CheckoutClient checkoutservice.Client
	OrderClient    orderservice.Client
	AuthClient     authservice.Client
	PaymentClient  paymentservice.Client
	once           sync.Once
	err            error
	registryAddr   string
	commonSuite    client.Option
)

func Init() {
	once.Do(func() {
		registryAddr = conf.GetConf().Hertz.RegistryAddress
		commonSuite = client.WithSuite(clientsuite.CommonClientSuite{
			RegistryAddr:       registryAddr,
			CurrentServiceName: conf.GetConf().Hertz.Service,
		})
		initUserClient()
		initProductClient()
		initCartClient()
		initCheckoutClient()
		initOrderClient()
		initAuthClient()
		initPaymentClient()
	})
}

func initUserClient() {
	UserClient, err = userservice.NewClient("user", commonSuite)
	frontendUtils.MustHandleError(err)
}

func initProductClient() {
	cbs := circuitbreak.NewCBSuite(func(ri rpcinfo.RPCInfo) string {
		return circuitbreak.RPCInfo2Key(ri)
	})
	cbs.UpdateServiceCBConfig("frontend/product/GetProduct",
		circuitbreak.CBConfig{Enable: true, ErrRate: 0.5, MinSample: 2},
	)

	ProductClient, err = productcatalogservice.NewClient("product",
		commonSuite,
		client.WithCircuitBreaker(cbs),
		client.WithFallback(
			fallback.NewFallbackPolicy(
				fallback.UnwrapHelper(
					func(ctx context.Context, req, resp interface{}, err error) (fbResp interface{}, fbErr error) {
						if err == nil {
							return resp, nil
						}
						methodName := rpcinfo.GetRPCInfo(ctx).To().Method()
						if methodName != "ListProducts" {
							return resp, err
						}
						return &product.ListProductsResp{
							Products: []*product.Product{
								{
									Price:       6.6,
									Id:          3,
									Picture:     "/static/image/t-shirt.jpeg",
									Name:        "T-Shirt",
									Description: "CloudWeGo T-Shirt",
								},
							},
						}, nil
					}),
			),
		))
	frontendUtils.MustHandleError(err)
}

func initCartClient() {
	CartClient, err = cartservice.NewClient("cart", commonSuite)
	frontendUtils.MustHandleError(err)
}

func initCheckoutClient() {
	CheckoutClient, err = checkoutservice.NewClient("checkout", commonSuite)
	frontendUtils.MustHandleError(err)
}

func initPaymentClient() {
	PaymentClient, err = paymentservice.NewClient("payment", commonSuite)
	frontendUtils.MustHandleError(err)
}

func initOrderClient() {
	OrderClient, err = orderservice.NewClient("order", commonSuite)
	frontendUtils.MustHandleError(err)
}

func initAuthClient() {
	AuthClient, err = authservice.NewClient("auth", commonSuite)
	frontendUtils.MustHandleError(err)
}
