package app

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/percona/pmm-agent/app/client"
	"github.com/percona/pmm-agent/app/config"
	"github.com/percona/pmm-agent/app/format"
	"github.com/percona/pmm-agent/app/server"
	"github.com/percona/pmm-agent/errs"
)

// App represents application.
type App struct {
	Client client.Client
	Server server.Server
	Config config.Config
	Format format.Format
}

func (a App) RunServer(ctx context.Context) error {
	errs := &errs.Safe{}
	wg := &sync.WaitGroup{}

	// Create context which cancels on termination signals.
	ctx = contextWithCancelOnSignal(ctx, syscall.SIGTERM, syscall.SIGINT)

	// Start all services.
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := a.Server.Serve(ctx); err != nil {
			errs.Add(err)
		}
	}()

	// Wait for all services to finish.
	wg.Wait()

	return errs.Err()

}

func contextWithCancelOnSignal(ctx context.Context, sig ...os.Signal) context.Context {
	ctx, cancel := context.WithCancel(ctx)

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, sig...)
	go func() {
		s := <-signals
		signal.Stop(signals)
		log.Printf("Got '%s (%d)' signal, shutting down...", s, s)
		cancel()
	}()

	return ctx
}
