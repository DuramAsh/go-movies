package models

type Movie struct {
	UUID   string  `json:"uuid" gorm:"primaryKey"`
	Title  string  `json:"title" binding:"required" gorm:"unique;not null"`
	Rating float64 `json:"rating" binding:"required" gorm:"not null"`
}

type UpdateMovie struct {
	UUID   string  `json:"uuid"`
	Title  string  `json:"title"`
	Rating float64 `json:"rating"`
}

type UserInput struct {
	Username string `json:"username" binding:"required" gorm:"primaryKey"`
	Password string `json:"password" binding:"required" gorm:"not null"`
}

type User struct {
	Username string `json:"username" binding:"required" gorm:"primaryKey"`
	Password []byte `json:"password" binding:"required" gorm:"not null"`
}
