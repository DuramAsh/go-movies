package models

import "gorm.io/gorm"

type Movie struct {
	gorm.Model
	Id     uint    `json:"id" gorm:"primaryKey"`
	Title  string  `json:"title"`
	Rating float64 `json:"rating"`
}
