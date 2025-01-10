package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func (cfg *Config) SetUser(username string) error {
	cfg.CurrentUserName = username

	err := write(*cfg)
	if err != nil {
		return fmt.Errorf("Failed to write : %w", err)
	}

	return nil
}

func Read() (Config, error) {
	var cfg Config

	filePath, err := getConfigFilePath()
	if err != nil {
		return cfg, fmt.Errorf("Failed to get config file path: %w", err)
	}

	file, err := os.Open(filePath)
	if err != nil {
		return cfg, fmt.Errorf("file not found: %w", err)
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&cfg)

	if err != nil {
		return cfg, fmt.Errorf("fail to decode json file: %w", err)
	}

	return cfg, nil

}

func write(cfg Config) error {
	getFilePath, err := getConfigFilePath()

	if err != nil {
		return fmt.Errorf("Failed to get config file path: %w", err)
	}

	file, err := os.Create(getFilePath)
	if err != nil {
		return fmt.Errorf("Failed to create file: %w", err)
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(cfg)

	if err != nil {
		return fmt.Errorf("Failed to encode struct to json: %w", err)
	}

	return nil

}

func getConfigFilePath() (string, error) {
	homePath, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("Home path is unknown: %w", err)
	}
	configPath := homePath + "/" + configFileName

	return configPath, nil

}
