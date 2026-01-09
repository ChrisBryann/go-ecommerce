package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/ChrisBryann/go-ecommerce/cmd/api"
	"github.com/ChrisBryann/go-ecommerce/config"
	"github.com/ChrisBryann/go-ecommerce/db"
)

func main() {

	db, err := db.NewPGXStorage(db.GenerateConnectionString(config.Envs.DBHost, config.Envs.DBPort, config.Envs.DBName, config.Envs.DBUser, config.Envs.DBPassword))

	if err != nil {
		log.Fatal(err)
	}

	initStorage(db)

	server := api.NewAPIServer(fmt.Sprintf(":%s", config.Envs.Port), db)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal((err))
	}

	log.Println("DB: Sucessfully connected!")
}
