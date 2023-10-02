package db

import (
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

// Runs up migrations to setup our database tables
func (d *Database) MigrateDB() error {
	fmt.Println("Migrating database")

	driver, err := postgres.WithInstance(d.Client.DB, &postgres.Config{})

	if err != nil {
		return fmt.Errorf("Could not create postgres driver: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance("file:///migrations", "postgres", driver)

	if err != nil {
		fmt.Println(err)
		return err
	}

	if err := m.Up(); err != nil {
		if !errors.Is(err, migrate.ErrNoChange) {
			return fmt.Errorf("Could not run up migrations: %w", err)
		}
	}

	fmt.Println("Sucessfully migrated database")

	return nil
}
