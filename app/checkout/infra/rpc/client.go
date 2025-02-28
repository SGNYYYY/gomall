package rpc

import (
	"sync"

	"github.com/SGNYYYY/gomall/app/checkout/conf"
	"github.com/SGNYYYY/gomall/common/clientsuite"
	"github.com/SGNYYYY/gomall/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/SGNYYYY/gomall/rpc_gen/kitex_gen/order/orderservice"
	"github.com/SGNYYYY/gomall/rpc_gen/kitex_gen/payment/paymentservice"
	"github.com/SGNYYYY/gomall/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/cloudwego/kitex/client"
)

var (
	CartClient    cartservice.Client
	ProductClient productcatalogservice.Client
	PaymentClient paymentservice.Client
	OrderClient   orderservice.Client
	once          sync.Once
	ServiceName   = conf.GetConf().Kitex.Service
	RegistryAddr  = conf.GetConf().Registry.RegistryAddress[0]
	commonSuite   client.Option
	err           error
)

func InitClient() {
	once.Do(func() {
		commonSuite = client.WithSuite(clientsuite.CommonClientSuite{
			CurrentServiceName: ServiceName,
			RegistryAddr:       RegistryAddr,
		})
		initCartClient()
		initProductClient()
		initOrderClient()
	})
}

func initCartClient() {
	CartClient, err = cartservice.NewClient("cart", commonSuite)
	if err != nil {
		panic(err)
	}
}

func initProductClient() {
	ProductClient, err = productcatalogservice.NewClient("product", commonSuite)
	if err != nil {
		panic(err)
	}
}

func initOrderClient() {
	OrderClient, err = orderservice.NewClient("order", commonSuite)
	if err != nil {
		panic(err)
	}
}
