package main

import (
	"errors"
	"flag"
	"github.com/anhgelus/go-anhgelus/data"
	"github.com/anhgelus/go-anhgelus/handler"
	slug "github.com/anhgelus/human-readable-slug"
	"github.com/gorilla/mux"
	"github.com/pelletier/go-toml/v2"
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"
)

var (
	verbose bool
	help    bool
)

func init() {
	flag.BoolVar(&verbose, "v", false, "Verbose mode")
	flag.BoolVar(&help, "h", false, "Show the help")
	flag.Parse()
}

func main() {
	if help || len(os.Args) < 2 {
		showHelp()
		return
	}
	if verbose {
		slog.SetLogLoggerLevel(slog.LevelDebug)
	}
	// loading configs
	slog.Info("Getting configs", "folder", "config")
	var err error
	data.Cfg, err = data.GetConfig()
	if err != nil {
		panic(err)
	}
	// handling commands
	command := os.Args[1]
	switch command {
	case "run":
		server()
	case "config":
		createConfig()
	default:
		showHelp()
		os.Exit(1)
	}
}

func server() {
	// launch server
	r := mux.NewRouter()
	r.HandleFunc("/{slug}", handler.Redirect)
	srv := &http.Server{
		Handler: r,
		Addr:    ":80",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	slog.Info("Starting http server")
	log.Fatal(srv.ListenAndServe())
}

func createConfig() {
	if len(os.Args) < 4 {
		slog.Debug("Path or url is empty")
		showHelp()
		os.Exit(2)
	}
	path := os.Args[2]
	url := os.Args[3]
	if data.Cfg.Has(url) {
		slog.Debug("url already used in config")
		os.Exit(3)
	}
	var b []byte
	var cfg data.Config
	if _, err := os.Stat("config/" + path); err == nil {
		by, err := os.ReadFile("config/" + path)
		if err != nil {
			panic(err)
		}
		err = toml.Unmarshal(by, &cfg)
		if err != nil {
			panic(err)
		}
		cfg.Links = append(
			cfg.Links,
			&data.LinkConfig{
				ID:   slug.GenerateSlug(uint64(time.Now().Unix()), 6),
				Link: url,
			},
		)
	} else if errors.Is(err, os.ErrNotExist) {
		cfg = data.Config{
			Links: []*data.LinkConfig{
				{
					ID:   slug.GenerateSlug(uint64(time.Now().Unix()), 6),
					Link: url,
				},
			},
		}
	} else {
		panic(err)
	}
	b, err := toml.Marshal(cfg)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("config/"+path, b, 754)
	if err != nil {
		panic(err)
	}
}

func showHelp() {
	println("Help of go-anhgelus")
	println("Command:")
	println("  - run -> start the http server")
	println(
		"  - config {path} {url} " +
			"-> create a new config (or edit the config) at the given path for the given url",
	)
	println("Flag:")
	println("  - -v -> Verbose mode")
	println("  - -h -> Show the help (this)")
}
