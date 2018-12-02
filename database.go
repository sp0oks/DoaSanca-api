package main

import (
    "os"

    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres" 
)

var db *gorm.DB

func getDB() (*gorm.DB, error) {
    var err error
    db, err = gorm.Open("postgres", os.Getenv("DATABASE_URL"))
    return db, err
}

func pingDB() error {
    err := db.DB().Ping()
    return err
}

func setupDB() error {
    err := pingDB()
    
    if err != nil {
        return err
    }
    if !db.HasTable(&User{}) {
        db.CreateTable(&User{})
        db.Create(&User{
                Name: "Gabriel Alves", 
                Email: "g4briel.4lves@gmail.com", 
                Coordinates: Coordinate{Latitude: -22.0027819, 
                                        Longitude:-47.8970543},
        })
    }
    if !db.HasTable(&Location{}) {
        db.CreateTable(&Location{})
        db.Create(&Location{
                Name: "Nave Sal da Terra",
                Type: "Brinquedo",
                Phone: 1633727823,
                Address: Address{Street: "R. Dep. Antonio Donato",
                                 Number: 428,
                                 Zipcode: 13573560},
            })
    }
    
    return nil
}

func getLocationByName(name string) Location {
    var location Location
    db.Where("name = ?", name).Find(&location)
    if db.RecordNotFound() {
        location.Name = ""
    }
    return location
}

func getLocations() Location {
    var locations Location
    db.Find(&locations)
    if db.RecordNotFound() {
        locations.Name = ""
    }
    return locations
}

func getUserByEmail(email string) User {
    var user User
    db.Where("email = ?", email).Find(&user)
    if db.RecordNotFound() {
        user.Name = ""
    }
    return user
}

func getUsers() User {
    var users User
    db.Find(&users)
    if db.RecordNotFound() {
        users.Name = ""
    }
    return users
}

