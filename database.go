package main

import (
    "os"
    "context"
    "database/sql"

    _ "github.com/jinzhu/gorm/dialects/postgres" 
)

func getDB() (*sql.DB, error) {
    db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
    return db, err
}

func pingDB (db *sql.DB) error {
    ctx := context.Background()
    err := db.PingContext(ctx)
    return err
}

func setupDB (db *sql.DB) error {
    return nil
}
