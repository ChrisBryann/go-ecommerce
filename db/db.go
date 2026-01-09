package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/ChrisBryann/go-ecommerce/config"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func NewPGXStorage(connString string) (*sql.DB, error) {
	db, err := sql.Open("pgx", connString)

	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}

func GenerateConnectionString(host string, port string, database string, username string, password string) string {
	return fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable", config.Envs.DBUser, config.Envs.DBPassword, config.Envs.DBHost, config.Envs.DBPort, config.Envs.DBName)
}
