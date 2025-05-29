package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	// TODO: Не используется Env поле.
	Env         string `yaml:"env"`
	StoragePath string `yaml:"storage_path"`
}

func GetConfig() *Config {
	// TODO: Хардкодить пусть к файлам конфигурации плохой тон. У тебя гибкие
	// файлы конфигурации, но лежать они могут только по этому пути. Надо получить
	// путь к файлу в main через os.Getenv() или аргументы приложения. Потом
	// уже этот файл можно прочитать.
	//
	// Например, файлом может быть несколько local.yaml, prod.yaml, dev.yaml и тд.
	envFile, err := os.ReadFile("./config/local.yaml")
	if err != nil {
		// TODO: Обработать ошибку нормально.
		// TODO: Использовать логгер.
		fmt.Println("Error reading config file")
	}

	var config Config
	err = yaml.Unmarshal(envFile, &config)
	if err != nil {
		// TODO: Обработать ошибку нормально.
		// TODO: Использовать логгер.
		fmt.Println("Error parsing config file")
	}

	return &config
}
