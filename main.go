package main

import (
	"flag"
	"github.com/anhgelus/go-anhgelus/data"
	"github.com/anhgelus/go-anhgelus/handler"
	"github.com/gorilla/mux"
	"log"
	"log/slog"
	"net/http"
	"time"
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
	var err error
	data.Cfg, err = data.GetConfig()
	if err != nil {
		panic(err)
	}
	// handling commands
	// launch server
	r := mux.NewRouter()
	r.HandleFunc("/{slug}", handler.Redirect)
	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:80",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
