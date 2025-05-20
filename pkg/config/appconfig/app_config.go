package appconfig

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/chazool/serendib_asia_service/pkg/config"
	lg "github.com/chazool/serendib_asia_service/pkg/log"
	"github.com/chazool/serendib_asia_service/pkg/web"
	"github.com/chazool/serendib_asia_service/pkg/web/middleware"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

const contextTimeout = 5 * time.Second

type service struct {
	_      struct{}
	App    *fiber.App
	Ctx    context.Context
	Cancel context.CancelFunc
}

// Start initializes the application and starts listening for requests.
func Start(routes func(*fiber.App)) {
	appConfig := config.GetConfig()
	app := web.SetupFiber(appConfig.ChildFiberProcessIdleTimeout)

	ctx, cancel := context.WithCancel(context.Background())

	defer shutdown(service{App: app, Ctx: ctx, Cancel: cancel})

	// add custom request middleware for panic recovery request, Request time & Request ID
	middleware.RequestMiddleware(app, appConfig.Pprofenabled)
	// Add cors Moddleware
	middleware.CorsMiddleware(app)

	// call routes
	routes(app)

	// fiber will listen from a defferent gorouting
	go func() {
		if err := app.Listen(":" + appConfig.SrvListenPort); err != nil {
			log.Panic(err)
		}
	}()

	c := make(chan os.Signal, 1)                    // create chanel to signal begin send
	signal.Notify(c, os.Interrupt, syscall.SIGTERM) // when an interrupt or termination signal is send, notif the channel

	<-c // this blocks the main thread until an interrupt is received
}

func shutdown(service service) {
	defer service.Cancel()

	// cleanly shutdown and flush telementry when the application exits
	defer func(ctx context.Context) {
		// do not make the application hang when it is shutdown
		_, cancel := context.WithTimeout(ctx, contextTimeout)
		defer cancel()
	}(service.Ctx)

	defer lg.Logger.Sync()

	err := web.Shutdown(service.App)
	if err != nil {
		lg.Logger.Fatal("Error during shudown", zap.Error(err))
	}
}
