package main

import (
	"github.com/airo507/GoProjectCore/internal/app"
	"github.com/airo507/GoProjectCore/internal/config"
	"log/slog"
	"os"
)

func main() {
	envConfig := config.GetConfig()
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	app.Run(envConfig, logger)
}
