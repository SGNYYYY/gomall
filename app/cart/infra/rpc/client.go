package rpc

import (
	"sync"

	"github.com/SGNYYYY/gomall/app/cart/conf"
	cartutils "github.com/SGNYYYY/gomall/app/cart/utils"
	"github.com/SGNYYYY/gomall/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
)

var (
	ProductClient productcatalogservice.Client
	once          sync.Once
)

func InitClient() {
	once.Do(func() {
		initProductClient()
	})
}

func initProductClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	cartutils.MustHandleError(err)
	ProductClient, err = productcatalogservice.NewClient("product", client.WithResolver(r))
	cartutils.MustHandleError(err)
}
