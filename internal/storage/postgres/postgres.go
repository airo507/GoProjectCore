package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	"log/slog"
)

func New(database string, logger *slog.Logger) (*sql.DB, error) {

	db, err := sql.Open("pgx", database)

	db.Ping()
	if err != nil {
		logger.Info("Error opening database %s\n", err)
		return nil, fmt.Errorf("%s", err)
	}

	return db, nil
}

func CheckDatabase(database string) error {
	db, err := sql.Open("pgx", database)
	if err != nil {
		slog.Info("Error opening database %s\n", err)
		return err
	}

	pingErr := db.Ping()
	if pingErr != nil {
		slog.Info("Error pinging database %s\n", pingErr)
		return pingErr
	}
	defer db.Close()
	return nil
}
