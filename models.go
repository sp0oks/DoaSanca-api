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
    Email       string      `json:"email"     binding:"required" gorm:"unique"`
    Latitude    float64     `json:"latitude"  binding:"required"`
    Longitude   float64     `json:"longitude" binding:"required"`
}
    
type Location struct {
    gorm.Model
    Name        string      `json:"name"    binding:"required" gorm:"unique"`
    Type        string      `json:"type"    binding:"required"`
    Email       string      `json:"email"   gorm:"default:''"`
    Phone       int64       `json:"phone"   binding:"required"`
    Street      string      `json:"street"  binding:"required"`
    Optional    string      `json:"opt"     gorm:"default:''"`
    Number      int64       `json:"number"  binding:"required"`
    Zipcode     int64       `json:"zipcode" binding:"required"`
    Comment     string      `json:"comment" gorm:"default:''"`
    PictureURL  string      `json:"pic_url" gorm:"default:''"`
}
