package main

import (
    "github.com/jinzhu/gorm"
)

type Coordinate struct {
    Latitude    float64     `json:"latitude"`
    Longitude   float64     `json:"longitude"`
}

type Address struct {
    Street      string      `json:"street"`
    Optional    string      `json:"opt"     gorm:"default:'null'"`
    Number      int64       `json:"number"`
    Zipcode     int64       `json:"zipcode"`
}

type User struct {
    gorm.Model
    Name        string      `json:"name"    binding:"required"`
    Email       string      `json:"email"   binding:"required"`
    Coordinates Coordinate  `json:"coords"`
}
    
type Location struct {
    gorm.Model
    Name        string      `json:"name"    binding:"required"`
    Type        string      `json:"type"    binding:"required"`
    Email       string      `json:"email"   gorm:"default:'null'"`
    Phone       int64       `json:"phone"   binding:"required"`
    Address     Address     `json:"address"`
    Comment     string      `json:"comment" gorm:"default:'null'"`
    PictureURL  string      `json:"pic_url" gorm:"default:'null'"`
}
