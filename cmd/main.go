package main

import (
	"database/sql"
	"log"
	"os"
	"os/signal"

	_ "github.com/lib/pq"
	"github.com/whatdislol/mobility-app/cmd/api"
	"github.com/whatdislol/mobility-app/config"
	"github.com/whatdislol/mobility-app/db"
)

func main() {
    db, err := db.NewPostgresStorage(config.Config{
        DBUser: config.Envs.DBUser,
        DBPassword: config.Envs.DBPassword,
        DBHost: config.Envs.DBHost,
        DBPort: config.Envs.DBPort,
        DBName: config.Envs.DBName,
    })

    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    initStorage(db)

    server := api.NewAPIServer(":8080", db)
    if err := server.Run(); err != nil {
        log.Fatal(err)
    }

    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt)
    go func() {
        <-c
        log.Println("Shutting down...")
        os.Exit(0)
    }()

}

func initStorage(db *sql.DB) {
    err := db.Ping()
    if err != nil {
        log.Fatal(err)
    }

    log.Println("DB successfully connected")
}