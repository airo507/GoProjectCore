package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func New(database string) (*sql.DB, error) {

	db, err := sql.Open("pgx", database)

	if err != nil {
		fmt.Printf("Error opening database %s\n", err)
		return nil, fmt.Errorf("%s", err)
	}

	query := []string{
		`CREATE TABLE IF NOT EXISTS users
		(
		    id SERIAL PRIMARY KEY,
		    login TEXT UNIQUE NOT NULL,
		    first_name TEXT NOT NULL,
		    last_name TEXT NOT NULL,
		    email TEXT NOT NULL,
		    password TEXT NOT NULL,
		    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);`,
		`CREATE TABLE IF NOT EXISTS posts
		(
		 	id SERIAL PRIMARY KEY,
		    author_id INTEGER NOT NULL,
		    body TEXT NOT NULL,
		    likes INTEGER,
		    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		    FOREIGN KEY (author_id) REFERENCES users (id)
		);`,
		`CREATE TABLE IF NOT EXISTS comments
		(
		 	id SERIAL PRIMARY KEY,
		    author_id INTEGER NOT NULL,
		    post_id INTEGER NOT NULL,
		    body TEXT NOT NULL,
		    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		    FOREIGN KEY (author_id) REFERENCES users (id),
		    FOREIGN KEY (post_id) REFERENCES posts (id)
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

	return db, nil
}
