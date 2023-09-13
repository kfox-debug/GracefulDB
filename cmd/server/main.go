package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"log/slog"

	"github.com/Kwynto/GracefulDB/internal/analyzers/sqlanalyzer"
	"github.com/Kwynto/GracefulDB/internal/base/basicsystem"
	"github.com/Kwynto/GracefulDB/internal/base/core"
	"github.com/Kwynto/GracefulDB/internal/config"
	"github.com/Kwynto/GracefulDB/internal/connectors/grpc"
	"github.com/Kwynto/GracefulDB/internal/connectors/rest"
	"github.com/Kwynto/GracefulDB/internal/connectors/socketconnector"
	"github.com/Kwynto/GracefulDB/internal/lib/helpers/loghelper"
	"github.com/Kwynto/GracefulDB/internal/manage/webmanage"
	"github.com/Kwynto/GracefulDB/pkg/lib/closer"
)

func runServer(ctx context.Context, cfg *config.Config) error {
	var closeProcs = &closer.Closer{}

	// TODO: Load the core of the system
	go core.Engine(cfg)
	closeProcs.Add(core.Shutdown) // Register a shutdown handler.

	// TODO: Run the basic command system
	go basicsystem.CommandSystem(cfg)
	closeProcs.Add(basicsystem.Shutdown) // Register a shutdown handler.

	// TODO: Start the language analyzer (SQL)
	go sqlanalyzer.Analyzer(cfg)
	closeProcs.Add(sqlanalyzer.Shutdown) // Register a shutdown handler.

	// TODO: Start Socket connector
	go socketconnector.Start(cfg)
	closeProcs.Add(socketconnector.Shutdown) // Register a shutdown handler.

	// TODO: Start REST API connector
	go rest.Start(cfg)
	closeProcs.Add(rest.Shutdown) // Register a shutdown handler.

	// TODO: Start gRPC connector
	go grpc.Start(cfg)
	closeProcs.Add(grpc.Shutdown) // Register a shutdown handler.

	// TODO: Start web-server for manage system
	go webmanage.Start(cfg)
	closeProcs.Add(webmanage.Shutdown) // Register a shutdown handler.

	// Waiting for a stop signal from the OS
	<-ctx.Done()
	slog.Warn("The shutdown process has started.")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), cfg.ShutdownTimeOut)
	defer cancel()

	if err := closeProcs.Close(shutdownCtx); err != nil {
		return fmt.Errorf("server shutdown: %v", err)
	}

	slog.Info("All processes are stopped.")

	return nil
}

func main() {
	// Init config
	configPath := os.Getenv("CONFIG_PATH")
	cfg := config.MustLoad(configPath)

	// Init logger: slog
	loghelper.Init(cfg)
	slog.Info("starting GracefulDB", slog.String("env", cfg.Env))
	slog.Debug("debug messages are enabled")

	// Signal tracking
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	if err := runServer(ctx, cfg); err != nil {
		slog.Error("An unexpected error occurred while the server was running.", slog.String("err", err.Error()))
	}

	slog.Info("GracefulDB has finished its work and will miss you.")
}
