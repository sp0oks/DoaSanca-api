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
    Optional    string      `json:"opt"`
    Number      int64       `json:"number"`
    Zipcode     int64       `json:"zipcode"`
}

type User struct {
    gorm.Model
    Name        string      `json:"name"      binding:"required"`
    Email       string      `json:"email"     binding:"required"`
    Latitude    float64     `json:"latitude"  binding:"required"`
    Longitude   float64     `json:"longitude" binding:"required"`
}
    
type Location struct {
    gorm.Model
    Name        string      `json:"name"`
    Type        string      `json:"type"`
    Email       string      `json:"email"   gorm:"default:'none'"`
    Phone       int64       `json:"phone"`
    Street      string      `json:"street"`
    Optional    string      `json:"opt"     gorm:"default:'none'"`
    Number      int64       `json:"number"`
    Zipcode     int64       `json:"zipcode"`
    Comment     string      `json:"comment" gorm:"default:'none'"`
    PictureURL  string      `json:"pic_url" gorm:"default:'none'"`
}
