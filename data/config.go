package data

import (
	"errors"
	slug "github.com/anhgelus/human-readable-slug"
	"github.com/pelletier/go-toml/v2"
	"log/slog"
	"os"
	"slices"
	"strings"
	"time"
)

type LinkConfig struct {
	ID   string `toml:"id"`
	Link string `toml:"link"`
}

type Config struct {
	Links []*LinkConfig `toml:"links"`
}

var Cfg *Config

func (c *Config) GetLink(slug string) string {
	for _, l := range c.Links {
		if l.ID == slug {
			return l.Link
		}
	}
	return ""
}

func (c *Config) GetLinkConfig(slug string) *LinkConfig {
	for _, l := range c.Links {
		if l.ID == slug {
			return l
		}
	}
	return nil
}

func (c *Config) Has(link string) bool {
	return slices.ContainsFunc(c.Links, func(l *LinkConfig) bool {
		return l.Link == link
	})
}

func (l *LinkConfig) GenerateID(path string) {
	l.ID = slug.GenerateSlug(uint64(time.Now().Unix()), 6)
	b, err := toml.Marshal(l)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile(path, b, 0754)
	if err != nil {
		panic(err)
	}
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
	return getConfigInDir(dir, "config/")
}

func getConfigInDir(dir []os.DirEntry, path string) (*Config, error) {
	cfg := Config{}
	for _, e := range dir {
		if e.IsDir() {
			d, err := os.ReadDir(path + e.Name())
			if err != nil {
				return nil, err
			}
			conf, err := getConfigInDir(d, path+e.Name()+"/")
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
		conf := Config{}
		b, err := os.ReadFile(path + e.Name())
		if err != nil {
			return nil, err
		}
		err = toml.Unmarshal(b, &conf)
		if err != nil {
			return nil, err
		}
		for _, l := range conf.Links {
			if l.ID == "" {
				l.GenerateID(path + e.Name())
				slog.Info("Generated slug", "link", l.Link, "slug", l.ID)
			}
		}
		cfg.Links = append(cfg.Links, conf.Links...)
	}
	return &cfg, nil
}
