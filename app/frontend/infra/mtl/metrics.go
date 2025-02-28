package mtl

import (
	"context"
	"net"
	"net/http"

	"github.com/SGNYYYY/gomall/app/frontend/conf"
	"github.com/SGNYYYY/gomall/common/utils"
	"github.com/cloudwego/hertz/pkg/app/server/registry"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/route"
	consulapi "github.com/hashicorp/consul/api"
	"github.com/hertz-contrib/registry/consul"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var Registry *prometheus.Registry

func InitMetric() route.CtxCallback {
	ServiceName := conf.GetConf().Hertz.Service
	Registry = prometheus.NewRegistry()
	Registry.MustRegister(collectors.NewGoCollector())
	Registry.MustRegister(collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}))
	config := consulapi.DefaultConfig()
	config.Address = conf.GetConf().Hertz.RegistryAddress
	consulClient, _ := consulapi.NewClient(config)
	r := consul.NewConsulRegister(consulClient)

	localIp := utils.MustGetLocalIPv4()
	ip, err := net.ResolveTCPAddr("tcp", localIp+conf.GetConf().Hertz.MetricsPort)
	if err != nil {
		hlog.Error(err)
	}
	registryInfo := &registry.Info{Addr: ip, ServiceName: "prometheus", Weight: 1, Tags: map[string]string{"service": ServiceName}}
	err = r.Register(registryInfo)
	if err != nil {
		hlog.Error(err)
	}

	http.Handle("/metrics", promhttp.HandlerFor(Registry, promhttp.HandlerOpts{}))
	go http.ListenAndServe(conf.GetConf().Hertz.MetricsPort, nil) //nolint:errcheck
	return func(ctx context.Context) {
		r.Deregister(registryInfo) //nolint:errcheck
	}
}
