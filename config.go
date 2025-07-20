package main

import (
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

type Config struct {
	MusingsFile string `toml:"musings_file"`
	CharLimit   int    `toml:"char_limit"`
	ExportPath  string `toml:"export_path"`
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
		config.CharLimit = 128
		config.ExportPath = filepath.Join(homeDir, "mlog", "musings.md")
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
