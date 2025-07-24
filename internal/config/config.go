package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

type EnvConfig struct {
	Dsn         string
	Host        string
	SecretKey   string
	KafkaConfig KafkaConfig
}

type KafkaConfig struct {
	KafkaAddr string
	Topic     string
	GroupId   string
	Partition int64
}

func GetConfig() *EnvConfig {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println("Error reading env file")
	}

	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	secretKey := os.Getenv("SECRET_KEY")
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	kafkaConfig := KafkaConfig{
		KafkaAddr: os.Getenv("KAFKA_ADDR"),
		Topic:     os.Getenv("KAFKA_TOPIC"),
		GroupId:   os.Getenv("KAFKA_GROUP_ID"),
		Partition: 0,
	}

	config := &EnvConfig{
		Dsn:         fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPassword, dbHost, dbPort, dbName),
		Host:        fmt.Sprintf("%s:%s", host, port),
		SecretKey:   secretKey,
		KafkaConfig: kafkaConfig,
	}

	return config
}
