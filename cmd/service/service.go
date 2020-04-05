package service

import (
	"os"
	"os/signal"
	"syscall"

	platform "bitbucket.org/celsa/hermes/env"

	"github.com/ea3hsp/book/pkg/api"
	"github.com/ea3hsp/book/pkg/render"
	"github.com/ea3hsp/book/pkg/store"
	"github.com/go-kit/kit/log"
)

const (
	// Template paths definitions
	includeTemplates = "./public/templates/"
	layoutTemplates  = "./public/templates/layout/"
	dbName           = "books"
	// Default config definitions
	defProcName = "book"
	defAPIPort  = "13000"
	defAPIPass  = "P4sZw0rD"
	// Environment variable names
	envProcName = "PROCESS_NAME"
	envAPIPort  = "API_PORT"
	envAPIPass  = "API_PASS"
)

// config struct definition
type config struct {
	processName string
	apiPort     string
	apiPass     string
}

// Run main func
func Run() {
	// parse os args
	cfg := loadConfig()
	// Creates logger
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}
	// creates render
	render := render.New(layoutTemplates, includeTemplates)
	err := render.LoadTemplates()
	if err != nil {
		logger.Log("[error]", "failed templates loading", "msg", err.Error())
	}
	// banner
	logger.Log("[info]", "loading templates successfully")
	// create store
	store := store.NewSimDB(dbName, logger)
	// create api
	api := api.NewAPI(logger, cfg.apiPort, cfg.apiPass, render, store)
	api.Init()
	// wait until signal
	sig := WaitSignal()
	// exit banner
	logger.Log("[Info]", "Exit", "signal", sig.String())
}

// WaitSignal catching exit signal
func WaitSignal() os.Signal {
	ch := make(chan os.Signal, 2)
	signal.Notify(
		ch,
		syscall.SIGINT,
		syscall.SIGQUIT,
		syscall.SIGTERM,
	)
	for {
		sig := <-ch
		switch sig {
		case syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM:
			return sig
		}
	}
}

// load config parameters
func loadConfig() *config {
	return &config{
		processName: platform.Env(envProcName, defProcName),
		apiPort:     platform.Env(envAPIPort, defAPIPort),
	}
}
