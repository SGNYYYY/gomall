package rpc

import (
	"sync"

	"github.com/SGNYYYY/gomall/app/frontend/conf"
	frontendUtils "github.com/SGNYYYY/gomall/app/frontend/utils"
	"github.com/SGNYYYY/gomall/rpc_gen/kitex_gen/auth/authservice"
	"github.com/SGNYYYY/gomall/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/SGNYYYY/gomall/rpc_gen/kitex_gen/checkout/checkoutservice"
	"github.com/SGNYYYY/gomall/rpc_gen/kitex_gen/order/orderservice"
	"github.com/SGNYYYY/gomall/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/SGNYYYY/gomall/rpc_gen/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
)

var (
	UserClient     userservice.Client
	ProductClient  productcatalogservice.Client
	CartClient     cartservice.Client
	ChechoutClient checkoutservice.Client
	OrderClient    orderservice.Client
	AuthClient     authservice.Client
	once           sync.Once
)

func Init() {
	once.Do(func() {
		initUserClient()
		initProductClient()
		initCartClient()
		initCheckoutClient()
		initOrderClient()
		initAuthClient()
	})
}

func initUserClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddress)
	frontendUtils.MustHandleError(err)
	UserClient, err = userservice.NewClient("user", client.WithResolver(r))
	frontendUtils.MustHandleError(err)
}

func initProductClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddress)
	frontendUtils.MustHandleError(err)
	ProductClient, err = productcatalogservice.NewClient("product", client.WithResolver(r))
	frontendUtils.MustHandleError(err)
}

func initCartClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddress)
	frontendUtils.MustHandleError(err)
	CartClient, err = cartservice.NewClient("cart", client.WithResolver(r))
	frontendUtils.MustHandleError(err)
}

func initCheckoutClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddress)
	frontendUtils.MustHandleError(err)
	ChechoutClient, err = checkoutservice.NewClient("checkout", client.WithResolver(r))
	frontendUtils.MustHandleError(err)
}

func initOrderClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddress)
	frontendUtils.MustHandleError(err)
	OrderClient, err = orderservice.NewClient("order", client.WithResolver(r))
	frontendUtils.MustHandleError(err)
}

func initAuthClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddress)
	frontendUtils.MustHandleError(err)
	AuthClient, err = authservice.NewClient("auth", client.WithResolver(r))
	frontendUtils.MustHandleError(err)
}
