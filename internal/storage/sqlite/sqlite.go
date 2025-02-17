package sqlite

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

type Storage struct {
	db *sql.DB
}

func CreateTables(database string) (*Storage, error) {

	db, err := sql.Open("sqlite3", database)

	if err != nil {
		fmt.Println("Error opening database %s", err)
		return nil, fmt.Errorf("%s", err)
	}
	defer db.Close()
	//transaction, err := db.Begin()

	query := []string{
		`CREATE TABLE IF NOT EXISTS user
		(
		    id INTEGER PRIMARY KEY AUTOINCREMENT,
		    login STRING NOT NULL,
		    first_name STRING NOT NULL,
		    last_name STRING NOT NULL,
		    email STRING NOT NULL,
		    password STRING NOT NULL,
		    created_at TIMESTAMP NOT NULL,
		    updated_at TIMESTAMP NOT NULL
		);`,
		`CREATE TABLE IF NOT EXISTS post
		(
		 	id INTEGER PRIMARY KEY AUTOINCREMENT,
		    author_id STRING NOT NULL,
		    body STRING NOT NULL,
		    likes INTEGER,
		    created_at TIMESTAMP NOT NULL,
		    updated_at TIMESTAMP NOT NULL 
		);`,
		`CREATE TABLE IF NOT EXISTS comment
		(
		 	id INTEGER PRIMARY KEY AUTOINCREMENT,
		    author_id STRING NOT NULL,
		    post_id STRING NOT NULL,
		    body STRING NOT NULL,
		    created_at TIMESTAMP NOT NULL,
		    updated_at TIMESTAMP NOT NULL
		);`,
	}

	for _, stmt := range query {
		_, err := db.Exec(stmt)
		if err != nil {
			fmt.Println("Error creating table: %s", err)
			return nil, fmt.Errorf("Error preparing statement %s", stmt)
		}
		fmt.Println("Tables created!")
	}

	return &Storage{db: db}, nil
}
