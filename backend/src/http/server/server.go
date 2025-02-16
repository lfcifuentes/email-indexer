package server

import (
	"context"
	"errors"
	"log"
	"log/slog"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/lfcifuentes/email-indexer/backend/src/core"
	"github.com/lfcifuentes/email-indexer/backend/src/http/router"
	"github.com/lfcifuentes/email-indexer/backend/src/services/email"
	"github.com/lfcifuentes/email-indexer/backend/src/services/zinc"
)

// StartServer starts the server
// It initializes the zinc service, email services and the router
// It starts the server and listens for incoming connections
// It handles graceful shutdown by waiting for active connections to finish before stopping the server
func StartServer() {
	// load app config
	config := core.BuildConfigFromEnv()

	// load app logger
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	// zinc service
	slog.Info("Initializing zinc", "zinc_server", config.ZincSearchAPIURL)
	zinc_server := zinc.NewZincService(
		config.ZincSearchAPIURL,
		config.ZincSearchAuthUser,
		config.ZincSearchAuthPass,
	)

	slog.Info("Pinging zinc services")
	err := zinc_server.Ping()
	if err != nil {
		slog.Error(err.Error(), "event", "zinc_connection_failed")
		// Exit the application if Zinc service is unavailable.
		os.Exit(1)
	} else {
		slog.Info("Zinc connection succeeded", "event", "zinc_connection_succeeded")
	}

	// email service
	slog.Info("Initializing email services")
	emailServices := email.NewEmailServices(config.EmailOutputDir)
	if err != nil {
		slog.Error(err.Error(), "event", "zinc_connection_failed")
		// Exit the application if Zinc service is unavailable.
		os.Exit(1)
	}

	// start server
	slog.Info("Starting server", "event", "starting_server", "port", config.HttpPort)
	r := router.NewRouter(
		config.HttpPort,
		zinc_server,
		emailServices,
	)

	// Start pprof server
	go func() {
		slog.Info("Starting pprof server", "port", "6060")
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	server := http.Server{
		Addr:    "0.0.0.0:" + config.HttpPort,
		Handler: r,
	}

	err = listenWithGracefulShutdown(config.ServerTearDownTimeOut, &server)
	if err != nil {
		slog.Error(err.Error(), "event", "server_listeners_failed")
	}
}

// listenWithGracefulShutdown starts the HTTP server and listens for incoming connections.
// It handles graceful shutdown by waiting for active connections to finish before stopping the server.
// In other words, when the system receives a termination signal (SIGINT or SIGTERM), it waits a specified
// timeout period to allow ongoing requests to complete before shutting down the server.
//
// Usage example:
//
//	srv := &http.Server{
//	    Addr:    ":3000",
//	    Handler: myRouter,
//	}
//	if err := listenWithGracefulShutdown(timeout, srv); err != nil {
//	    log.Fatalf("Server error: %v", err)
//	}
func listenWithGracefulShutdown(timeout time.Duration, srv *http.Server) error {
	shutdownCompleted := make(chan struct{})

	go func() {
		defer close(shutdownCompleted)

		osInterruptSignal := make(chan os.Signal, 1)
		signal.Notify(osInterruptSignal, syscall.SIGTERM, syscall.SIGINT)

		<-osInterruptSignal

		ctxTimeout, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		slog.Info("Shutting down server", "event", "server_shutting_down")

		err := srv.Shutdown(ctxTimeout)
		if err != nil {
			log.Fatal("server_listeners_shutdown_failed")
		}
	}()

	slog.Info("Server started", "event", "server_started")
	if err := srv.ListenAndServe(); errors.Is(err, http.ErrServerClosed) {
		return err
	}

	// ListenAndServe returns immediately after Shutdown is invoked (see doc)
	// This means we have to wait right after it returns to give enough time to the Shutdown method
	// to wait for all connections to close down gracefully
	<-shutdownCompleted

	slog.Info("Server shutdown completed", "event", "server_shutdown_completed")

	return nil
}
