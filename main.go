package main

import (
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
	path    string
	url     string
)

func init() {
	flag.BoolVar(&verbose, "v", false, "Verbose mode")
	flag.BoolVar(&help, "h", false, "Show the help")
	flag.StringVar(&path, "path", "", "Set the path")
	flag.StringVar(&url, "url", "", "Set the url")
}

func main() {
	flag.Parse()
	if help || len(os.Args) < 2 {
		showHelp()
		return
	}
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
	command := os.Args[1]
	switch command {
	case "run":
		server()
	case "create":
		createConfig()
	default:
		showHelp()
	}
}

func server() {
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
	slog.Info("Starting http server")
	log.Fatal(srv.ListenAndServe())
}

func createConfig() {
	if len(path) == 0 || len(url) == 0 {
		showHelp()
		return
	}
	cfg := data.Config{
		Links: []*data.LinkConfig{
			{
				ID:   slug.GenerateSlug(time.Now().Unix(), 6),
				Link: url,
			},
		},
	}
	b, err := toml.Marshal(cfg)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile(path, b, 754)
	if err != nil {
		panic(err)
	}
}

func showHelp() {
	println("Help of go-anhgelus")
	println("Command:")
	println("  - run -> start the http server")
	println("  - config --path {path} --url {url} -> create a new config at the given path for the given url")
	println("Flag:")
	println("  - -v -> Verbose mode")
	println("  - -h -> Show the help (this)")
}
