package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"golang.org/x/sync/errgroup"

	"github.com/danekin/aviasales-tz/config"
	"github.com/danekin/aviasales-tz/internal/app/dictionary"
	"github.com/danekin/aviasales-tz/internal/app/handlers"
	"github.com/danekin/aviasales-tz/internal/app/usecase/anagram"
	"github.com/danekin/aviasales-tz/internal/app/usecase/load"
)

const (
	FailedParseConfigCode = iota + 1
	FailedStartServerCode
	FailedStopServerCode
)

func main() {
	cfg, err := config.Parse()
	if err != nil {
		os.Exit(FailedParseConfigCode)
	}

	errGroup, ctx := errgroup.WithContext(context.Background())
	cancelCh := make(chan os.Signal, 1)

	signal.Notify(cancelCh, syscall.SIGINT, syscall.SIGTERM)

	dict := dictionary.NewCachedDictionary()

	srvRouter := router.New()
	srvRouter.POST("/load", handlers.NewLoadDictionary(load.NewLoadDictionaryUsecase(dict)).Handle)
	srvRouter.GET("/get", handlers.NewCheckWordHandler(anagram.NewFindAnagramUsecase(dict)).Handle)

	srv := fasthttp.Server{
		Handler: srvRouter.Handler,
	}

	errGroup.Go(func() error {
		return srv.ListenAndServe(fmt.Sprintf(":%v", cfg.HTTPServerConfig.Port))
	})

	select {
	case <-ctx.Done():
		fmt.Printf("failed to start http server: %v", errGroup.Wait())
		os.Exit(FailedStartServerCode)
	case <-cancelCh:
		fmt.Printf("got exit signal, stopping application\n")

		break
	}

	if err := srv.Shutdown(); err != nil {
		fmt.Printf("failed to stop http server: %v", err)
		os.Exit(FailedStopServerCode)
	}

	if err := errGroup.Wait(); err != nil {
		fmt.Printf("failed to stop stop application: %v", err)
		os.Exit(FailedStopServerCode)
	}

	fmt.Printf("application successfully stopped\n")
}
