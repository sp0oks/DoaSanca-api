package main

import (
    "github.com/jinzhu/gorm"
)

type Coordinate struct {
    Latitude    float64     `json:"latitude"  binding:"required"`
    Longitude   float64     `json:"longitude" binding:"required"`
}

type Address struct {
    Street      string      `json:"street"  binding:"required"`
    Optional    string      `json:"opt"     gorm:"default:'null'"`
    Number      int64       `json:"number"  binding:"required"`
    Zipcode     int64       `json:"zipcode" binding:"required"`
}

type User struct {
    gorm.Model
    Name        string      `json:"name"    binding:"required"`
    Email       string      `json:"email"   binding:"required"`
    Coordinates Coordinate  `json:"coords"  binding:"required"`
}
    
type Location struct {
    gorm.Model
    Name        string      `json:"name"    binding:"required"`
    Type        string      `json:"type"    binding:"required"`
    Email       string      `json:"email"   gorm:"default:'null'"`
    Phone       int64       `json:"phone"   binding:"required"`
    Address     Address     `json:"address" binding:"required"`
    Comment     string      `json:"comment" gorm:"default:'null'"`
    PictureURL  string      `json:"pic_url" gorm:"default:'null'"`
}
