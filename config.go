package main

import (
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

type Config struct {
	MusingsFile string `toml:"musings_file"`
}

func loadConfig() (Config, error) {
	var config Config
	configDir, err := os.UserConfigDir()
	if err != nil {
		return config, err
	}

	configPath := filepath.Join(configDir, "mlog", "config.toml")

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		// Create a default config if it doesn't exist
		homeDir, _ := os.UserHomeDir()
		config.MusingsFile = filepath.Join(homeDir, "mlog", "musings.ndjson")
		if err := os.MkdirAll(filepath.Dir(configPath), 0755); err != nil {
			return config, err
		}

		f, err := os.Create(configPath)
		if err != nil {
			return config, err
		}
		defer f.Close()
		return config, toml.NewEncoder(f).Encode(config)
	}

	if _, err := toml.DecodeFile(configPath, &config); err != nil {
		return config, err
	}
	return config, nil
}
