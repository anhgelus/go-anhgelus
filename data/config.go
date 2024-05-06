package data

import (
	"errors"
	"log/slog"
	"os"
	"strings"
)

type LinkConfig struct {
	ID   string `toml:"id"`
	Link string `toml:"link"`
}

type Config struct {
	Links []*LinkConfig `toml:"links"`
}

func GetConfig() (*Config, error) {
	dir, err := os.ReadDir("config")
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			slog.Info("Config folder does not exist. Creating a new one.")
			err = os.Mkdir("config", 0754)
			if err != nil {
				return nil, err
			}
			return &Config{}, nil
		}
		return nil, err
	}
	return getConfigInDir(dir, "")
}

func getConfigInDir(dir []os.DirEntry, path string) (*Config, error) {
	cfg := Config{}
	for _, e := range dir {
		if e.IsDir() {
			d, err := os.ReadDir(path + e.Name())
			if err != nil {
				return nil, err
			}
			conf, err := getConfigInDir(d, path+e.Name())
			if err != nil {
				return nil, err
			}
			cfg.Links = append(cfg.Links, conf.Links...)
			continue
		}
		if !strings.HasSuffix(e.Name(), ".toml") {
			slog.Debug("Skipping file", "name", e.Name())
			continue
		}
		// parse toml file
	}
	return &cfg, nil
}
