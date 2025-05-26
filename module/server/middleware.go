package server

import "go.uber.org/fx"

var middleware = fx.Module("middleware",
	fx.Invoke(
		configureLoggerMiddleware,
		configureOtelFiberConfig,
		configureZapLogger,
		configureCors,
		configureMonitor,
		configureHealthcheck,
		configureCompress,
		configureFavicon,
		configureDefaultPage,
		configureHelmet,
		configurePrometheus,
		configurePprof,
		configureRequestId,
		configureRecover,
		proxyMiddleware,
	),
)
