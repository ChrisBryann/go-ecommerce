package main

import (
	"log"
	"os"

	"github.com/ChrisBryann/go-ecommerce/config"
	"github.com/ChrisBryann/go-ecommerce/db"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/pgx"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	log.Printf("DB: host=%q port=%q user=%q db=%q",
		config.Envs.DBHost, config.Envs.DBPort, config.Envs.DBUser, config.Envs.DBName)

	db, err := db.NewPGXStorage(db.GenerateConnectionString(config.Envs.DBHost, config.Envs.DBPort, config.Envs.DBName, config.Envs.DBUser, config.Envs.DBPassword))

	if err != nil {
		log.Fatalf("failed setting new pgx storage: %v", err)
	}

	driver, err := pgx.WithInstance(db, &pgx.Config{})

	if err != nil {
		log.Fatalf("failed setting pgx driver: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://cmd/migrate/migrations",
		"postgres",
		driver,
	)

	if err != nil {
		log.Fatal(err)
	}

	cmd := os.Args[(len(os.Args) - 1)]

	if cmd == "up" {
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	}

	if cmd == "down" {
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	}

}
