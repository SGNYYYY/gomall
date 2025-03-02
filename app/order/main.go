package main

import (
	"context"
	"net"
	"time"

	"github.com/SGNYYYY/gomall/app/order/biz/dal"
	"github.com/SGNYYYY/gomall/app/order/conf"
	"github.com/SGNYYYY/gomall/app/order/infra/schedule"
	"github.com/SGNYYYY/gomall/common/mtl"
	"github.com/SGNYYYY/gomall/common/serversuite"
	"github.com/SGNYYYY/gomall/rpc_gen/kitex_gen/order/orderservice"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	"github.com/joho/godotenv"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	"github.com/reugn/go-quartz/job"
	"github.com/reugn/go-quartz/quartz"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	ServiceName  = conf.GetConf().Kitex.Service
	RegistryAddr = conf.GetConf().Registry.RegistryAddress[0]
)

func main() {
	_ = godotenv.Load()
	mtl.InitMetric(ServiceName, conf.GetConf().Kitex.MetricsPort, RegistryAddr)
	p := mtl.InitTracing(ServiceName)
	defer p.Shutdown(context.Background()) //nolint:errcheck

	dal.Init()
	opts := kitexInit()

	scheduler, _ := quartz.NewStdScheduler()
	scheduler.Start(context.Background())
	job_cancel := &schedule.OrderCancellationJob{}
	functionJob := job.NewFunctionJob(func(_ context.Context) (int, error) {
		job_cancel.Execute()
		return 1, nil
	})
	trigger := quartz.NewSimpleTrigger(1 * time.Minute)
	_ = scheduler.ScheduleJob(quartz.NewJobDetail(functionJob, quartz.NewJobKey("ScheduleCancelOrder")), trigger)

	svr := orderservice.NewServer(new(OrderServiceImpl), opts...)

	err := svr.Run()
	if err != nil {
		klog.Error(err.Error())
	}
}

func kitexInit() (opts []server.Option) {
	// address
	addr, err := net.ResolveTCPAddr("tcp", conf.GetConf().Kitex.Address)
	if err != nil {
		panic(err)
	}
	opts = append(opts, server.WithServiceAddr(addr), server.WithSuite(serversuite.CommonServerSuite{
		CurrentServiceName: ServiceName,
		RegistryAddr:       RegistryAddr,
	}))

	// klog
	logger := kitexlogrus.NewLogger()
	klog.SetLogger(logger)
	klog.SetLevel(conf.LogLevel())
	asyncWriter := &zapcore.BufferedWriteSyncer{
		WS: zapcore.AddSync(&lumberjack.Logger{
			Filename:   conf.GetConf().Kitex.LogFileName,
			MaxSize:    conf.GetConf().Kitex.LogMaxSize,
			MaxBackups: conf.GetConf().Kitex.LogMaxBackups,
			MaxAge:     conf.GetConf().Kitex.LogMaxAge,
		}),
		FlushInterval: time.Minute,
	}
	klog.SetOutput(asyncWriter)
	server.RegisterShutdownHook(func() {
		asyncWriter.Sync()
	})
	return
}
