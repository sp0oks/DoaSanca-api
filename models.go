package main

import (
    "time"

    "github.com/jinzhu/gorm"
)

type Coordinate struct {
    Latitude    float64     `json:"lat"     binding:"required"`
    Longitude   float64     `json:"lon"     binding:"required"`
}

type Address struct {
    Street      string      `json:"street"  binding:"required"`
    Optional    string      `json:"opt"`
    Number      int64       `json:"number"  binding:"required"`
    ZipCode     int64       `json:"zipcode" binding:"required"`
}

type User struct {
    gorm.Model
    Email       string      `json:"email"   binding:"required"`
    Name        string      `json:"name"    binding:"required"`
    Coordinates Coordinate  `json:"coords"  binding:"required"`
    Date        time.Time
}
    
type Location struct {
    gorm.Model
    Name        string      `json:"name"    binding:"required"`
    Type        string      `json:"type"    binding:"required"`
    Email       string      `json:"email"   binding:"required"`
    Phone       int64       `json:"phone"   binding:"required"`
    Address     Address     `json:"address" binding:"required"`
    Comment     string      `json:"comment"`
    Date        time.Time   
}
