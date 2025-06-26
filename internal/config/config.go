package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type LogConfig struct {
	ID   string `json:"id"`
	Path string `json:"path"`
	Type string `json:"type"`
}

func LoadConfig(path string) ([]LogConfig, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("impossible d'ouvrir le fichier de configuration: %w", err)
	}
	defer file.Close()

	var configs []LogConfig
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&configs); err != nil {
		return nil, fmt.Errorf("erreur de décodage JSON: %w", err)
	}
	return configs, nil
}
