package main

import (
	"flag"
	"github.com/anhgelus/go-anhgelus/data"
	"log/slog"
)

var (
	verbose bool
)

func init() {
	flag.BoolVar(&verbose, "v", false, "Verbose mode")
}

func main() {
	flag.Parse()
	if verbose {
		slog.SetLogLoggerLevel(slog.LevelDebug)
	}
	slog.Info("Starting go-anhgelus")
	// loading configs
	slog.Info("Getting configs", "folder", "config")
	cfg, err := data.GetConfig()
	if err != nil {
		panic(err)
	}
	// handling commands
}
