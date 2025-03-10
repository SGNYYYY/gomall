// Code generated by hertz generator.

package main

import (
	"context"
	"os"
	"time"

	"github.com/SGNYYYY/gomall/app/frontend/biz/router"
	"github.com/SGNYYYY/gomall/app/frontend/conf"
	"github.com/SGNYYYY/gomall/app/frontend/infra/mtl"
	"github.com/SGNYYYY/gomall/app/frontend/infra/rpc"
	"github.com/SGNYYYY/gomall/app/frontend/middleware"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/middlewares/server/recovery"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/cors"
	"github.com/hertz-contrib/gzip"
	"github.com/hertz-contrib/logger/accesslog"
	hertzlogrus "github.com/hertz-contrib/logger/logrus"
	hertzprom "github.com/hertz-contrib/monitor-prometheus"
	hertzobslogrus "github.com/hertz-contrib/obs-opentelemetry/logging/logrus"
	hertztracing "github.com/hertz-contrib/obs-opentelemetry/tracing"
	"github.com/hertz-contrib/pprof"
	"github.com/hertz-contrib/sessions"
	"github.com/hertz-contrib/sessions/redis"
	"github.com/joho/godotenv"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	ServiceName  = conf.GetConf().Hertz.Service
	MetricsPort  = conf.GetConf().Hertz.MetricsPort
	RegistryAddr = conf.GetConf().Hertz.RegistryAddress
)

func main() {
	// init dal
	// dal.Init()
	_ = godotenv.Load()
	mtl.InitMetric()
	p := mtl.InitTracing(ServiceName)
	defer p.Shutdown(context.Background()) //nolint:errcheck
	address := conf.GetConf().Hertz.Address
	rpc.Init()

	tracer, cfg := hertztracing.NewServerTracer()

	h := server.New(server.WithHostPorts(address), server.WithTracer(
		hertzprom.NewServerTracer(
			"",
			"",
			hertzprom.WithRegistry(mtl.Registry),
			hertzprom.WithDisableServer(true),
		),
	),
		tracer,
	)
	h.Use(hertztracing.ServerMiddleware(cfg))

	registerMiddleware(h)

	// add a ping route to test
	h.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
		ctx.JSON(consts.StatusOK, utils.H{"ping": "pong"})
	})

	router.GeneratedRegister(h)
	h.LoadHTMLGlob("template/*")
	h.Static("/static", "./")

	h.GET("/sign-in", func(ctx context.Context, c *app.RequestContext) {
		data := utils.H{
			"title": "Sign In",
			"next":  string(c.GetHeader("Referer")),
		}
		c.HTML(consts.StatusOK, "sign-in", data)
	})

	h.GET("/sign-up", func(ctx context.Context, c *app.RequestContext) {
		c.HTML(consts.StatusOK, "sign-up", utils.H{"title": "Sign Up"})
	})
	h.GET("/admin", func(ctx context.Context, c *app.RequestContext) {
		c.HTML(consts.StatusOK, "admin-login", utils.H{"title": "Admin"})
	})

	h.GET("/about", func(ctx context.Context, c *app.RequestContext) {
		hlog.CtxInfof(ctx, "CloudWeGo shop about page")
		c.HTML(consts.StatusOK, "about", utils.H{"title": "About"})
	})

	h.Spin()
}

func registerMiddleware(h *server.Hertz) {
	store, _ := redis.NewStore(10, "tcp", conf.GetConf().Redis.Address, "", []byte(os.Getenv("SEESION_SECRET")))
	h.Use(sessions.New("cloudwege-shop", store))

	// log
	logger := hertzobslogrus.NewLogger(hertzobslogrus.WithLogger(hertzlogrus.NewLogger().Logger()))
	// logger := hertzlogrus.NewLogger()
	hlog.SetLogger(logger)
	hlog.SetLevel(conf.LogLevel())
	var flushInterval time.Duration
	if os.Getenv("GO_ENV") == "online" {
		flushInterval = time.Minute
	} else {
		flushInterval = time.Second
	}
	asyncWriter := &zapcore.BufferedWriteSyncer{
		WS: zapcore.AddSync(&lumberjack.Logger{
			Filename:   conf.GetConf().Hertz.LogFileName,
			MaxSize:    conf.GetConf().Hertz.LogMaxSize,
			MaxBackups: conf.GetConf().Hertz.LogMaxBackups,
			MaxAge:     conf.GetConf().Hertz.LogMaxAge,
		}),
		FlushInterval: flushInterval,
	}
	hlog.SetOutput(asyncWriter)
	h.OnShutdown = append(h.OnShutdown, func(ctx context.Context) {
		asyncWriter.Sync()
	})

	// pprof
	if conf.GetConf().Hertz.EnablePprof {
		pprof.Register(h)
	}

	// gzip
	if conf.GetConf().Hertz.EnableGzip {
		h.Use(gzip.Gzip(gzip.DefaultCompression))
	}

	// access log
	if conf.GetConf().Hertz.EnableAccessLog {
		h.Use(accesslog.New())
	}

	// recovery
	h.Use(recovery.Recovery())

	// cores
	h.Use(cors.Default())

	middleware.Register(h)
}
