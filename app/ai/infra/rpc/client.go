package rpc

import (
	"sync"

	"github.com/SGNYYYY/gomall/app/checkout/conf"
	"github.com/SGNYYYY/gomall/common/clientsuite"
	"github.com/SGNYYYY/gomall/rpc_gen/kitex_gen/checkout/checkoutservice"
	"github.com/SGNYYYY/gomall/rpc_gen/kitex_gen/order/orderservice"
	"github.com/cloudwego/kitex/client"
)

var (
	OrderClient    orderservice.Client
	CheckoutClient checkoutservice.Client
	once           sync.Once
	ServiceName    = conf.GetConf().Kitex.Service
	RegistryAddr   = conf.GetConf().Registry.RegistryAddress[0]
	commonSuite    client.Option
	err            error
)

func InitClient() {
	once.Do(func() {
		commonSuite = client.WithSuite(clientsuite.CommonClientSuite{
			CurrentServiceName: ServiceName,
			RegistryAddr:       RegistryAddr,
		})
		initOrderClient()
		initCheckoutClient()
	})
}

func initOrderClient() {
	OrderClient, err = orderservice.NewClient("order", commonSuite)
	if err != nil {
		panic(err)
	}
}

func initCheckoutClient() {
	CheckoutClient, err = checkoutservice.NewClient("checkout", commonSuite)
	if err != nil {
		panic(err)
	}
}
