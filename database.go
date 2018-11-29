package main

import (
    "os"
    "log"
    "context"
    "time"
    "database/sql"

    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres" 
)

type Coordinate struct {
    Latitude    float64
    Longitude   float64
}

type Address struct {
    Street      string
    Optional    string
    Number      int64
    ZipCode     int64
}

type User struct {
    gorm.Model
    Email       string
    Name        string
    Coordinates Coordinate
    Date        time.Time
}
    
type Location struct {
    gorm.Model
    Name        string
    Type        string
    Email       string
    Telephone   int64
    Address     Address
    Comment     string
    Date        time.Time
}


func setupDB() {
    ctx := context.Background()
    db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
    
    if err != nil {
        log.Fatalf("Error while connecting database: %q", err)
    }
    
    err = db.PingContext(ctx)
    
    if err != nil {
        log.Fatalf("Error while checking connection to the database: %q", err)
    }
}
