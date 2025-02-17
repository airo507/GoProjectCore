package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Env         string `yaml:"env"`
	StoragePath string `yaml:"storage_path"`
}

func GetConfig() *Config {
	envFile, err := os.ReadFile("./config/local.yaml")
	if err != nil {
		fmt.Println("Error reading config file")
	}

	var config Config
	err = yaml.Unmarshal(envFile, &config)
	if err != nil {
		fmt.Println("Error parsing config file")
	}

	return &config
}
