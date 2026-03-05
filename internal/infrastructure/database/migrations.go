package database

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigrations(db *sql.DB) {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal("Failed to get current directory:", err)
	}

	migrationPath := filepath.Join(currentDir, "migrations")
	log.Printf("Migration path: %s", migrationPath)

	if _, err := os.Stat(migrationPath); os.IsNotExist(err) {
		log.Fatalf("Migrations folder not found at: %s", migrationPath)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal("Migration driver error:", err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://"+migrationPath,
		"postgres",
		driver,
	)
	if err != nil {
		log.Fatal("Migration init error:", err)
	}
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		log.Fatal("Migration failed:", err)
	}
	log.Println("Migration success")
}
