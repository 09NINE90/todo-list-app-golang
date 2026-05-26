package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	corelogger "github.com/09NINE90/todo-list-app-golang/internal/core/logger"
	corepoolconn "github.com/09NINE90/todo-list-app-golang/internal/core/repository/pool"
	corehttpmiddleware "github.com/09NINE90/todo-list-app-golang/internal/core/transport/http/middleware"
	corehttpserver "github.com/09NINE90/todo-list-app-golang/internal/core/transport/http/server"
	usersrepository "github.com/09NINE90/todo-list-app-golang/internal/features/users/repository"
	usersservice "github.com/09NINE90/todo-list-app-golang/internal/features/users/service"
	userstransporthttp "github.com/09NINE90/todo-list-app-golang/internal/features/users/transport/http"
	"go.uber.org/zap"
)

func main() {
	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT, syscall.SIGTERM,
	)
	defer cancel()

	logger, err := corelogger.NewLogger(corelogger.NewConfigMust())
	if err != nil {
		fmt.Println("failed to init application logger:", err)
		os.Exit(1)
	}
	defer logger.Close()

	logger.Debug("initializing connection pool")
	pool, err := corepoolconn.NewConnPool(
		ctx,
		corepoolconn.NewConfigMust(),
	)
	if err != nil {
		logger.Fatal("failed to init connection pool", zap.Error(err))
	}
	defer pool.Close()

	logger.Debug("initializing feature", zap.String("feature", "users"))
	usersRepository := usersrepository.NewUsersRepository(pool)
	usersService := usersservice.NewUsersService(usersRepository)
	usersTransportHTTP := userstransporthttp.NewUsersHttpHandler(usersService)

	logger.Debug("initializing HTTP server")
	httpServer := corehttpserver.NewHTTPServer(
		corehttpserver.NewConfigMust(),
		logger,
		corehttpmiddleware.RequestID(),
		corehttpmiddleware.Logger(logger),
		corehttpmiddleware.Panic(),
		corehttpmiddleware.Trace(),
	)
	apiVersionRouter := corehttpserver.NewApiVersionRouter(corehttpserver.ApiVersionV1)
	apiVersionRouter.RegisterRouters(usersTransportHTTP.Routers()...)
	httpServer.RegisterAPIRouters(apiVersionRouter)

	if err := httpServer.Run(ctx); err != nil {
		logger.Error("HTTP server run error", zap.Error(err))
	}
}
