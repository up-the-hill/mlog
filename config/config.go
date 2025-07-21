package config

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

type Config struct {
	MusingsFile string `toml:"musings_file"`
	ExportPath  string `toml:"export_path"`
	CharLimit   int    `toml:"char_limit"`
}

var homeDir, _ = os.UserHomeDir()

var defaultConfig Config = Config{
	filepath.Join(homeDir, "mlog", "musings.ndjson"),
	filepath.Join(homeDir, "mlog", "musings.md"),
	128,
}

func LoadConfig() Config {
	var config Config
	configDir, _ := os.UserConfigDir()

	configPath := filepath.Join(configDir, "mlog", "config.toml")

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return defaultConfig
	}

	toml.DecodeFile(configPath, &config)
	return config
}

func CreateConfig() error {
	var config Config
	configDir, _ := os.UserConfigDir()
	configPath := filepath.Join(configDir, "mlog", "config.toml")

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		// Create a default config if it doesn't exist
		homeDir, _ := os.UserHomeDir()
		config.MusingsFile = filepath.Join(homeDir, "mlog", "musings.ndjson")
		config.ExportPath = filepath.Join(homeDir, "mlog", "musings.md")
		config.CharLimit = 128
		os.MkdirAll(filepath.Dir(configPath), 0755)

		f, _ := os.Create(configPath)
		defer f.Close()
		toml.NewEncoder(f).Encode(config)
	} else {
		return errors.New("Config file already exists")
	}
	return nil
}
