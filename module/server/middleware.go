package server

import "go.uber.org/fx"

var middleware = fx.Module("middleware",
	fx.Invoke(
		configureMonitor,
		proxyMiddleware,
		configureLoggerMiddleware,
		configureOtelFiberConfig,
		configureZapLogger,
		configureCors,
		configureHealthcheck,
		configureCompress,
		configureFavicon,
		configureDefaultPage,
		configureHelmet,
		configurePrometheus,
		configurePprof,
		configureRequestId,
		configureRecover,
	),
)
