package models

import "gorm.io/gorm"

type Movie struct {
	gorm.Model
	Title  string  `json:"title"`
	Rating float64 `json:"rating"`
}

type APIMovie struct {
	Title  string  `json:"title"`
	Rating float64 `json:"rating"`
}
