package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Database struct {
	Client *sqlx.DB
}

// Creates a connection to an SQL database.
// The pointer to Database struct it returns contains a Client field which is used to send queries.
func NewDatabase() (*Database, error) {
	// environment vars are loaded from docker-compose.yaml

	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_TABLE"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("SSL_MODE"),
	)

	dbConn, err := sqlx.Connect("postgres", connectionString)

	if err != nil {
		return &Database{}, fmt.Errorf("error connecting to database: %w", err)
	}

	return &Database{Client: dbConn}, nil
}

// Checks if the database has a healthy connection
func (d *Database) Ping(ctx context.Context) error {
	return d.Client.DB.PingContext(ctx)
}
