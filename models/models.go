package models

type Movie struct {
	UUID      string  `json:"uuid" gorm:"primaryKey"`
	Title     string  `json:"title" binding:"required" gorm:"unique;not null"`
	Rating    float64 `json:"rating" binding:"required" gorm:"not null"`
	Year      uint    `json:"year" binding:"required" gorm:"not null"`
	PosterURL string  `json:"posterURL"`
}

type MovieInput struct {
	UUID      string  `json:"uuid"`
	Title     string  `json:"title"`
	Rating    float64 `json:"rating"`
	PosterURL string  `json:"posterURL"`
	Year      uint    `json:"year"`
}

type User struct {
	Username string  `json:"username" binding:"required" gorm:"primaryKey"`
	Password []byte  `json:"password" binding:"required" gorm:"not null"`
	Movies   []Movie `gorm:"many2many:user_movie_thoughts"`
}

type UserInput struct {
	Username string `json:"username" binding:"required" gorm:"primaryKey"`
	Password string `json:"password" binding:"required" gorm:"not null"`
}

type UserMovieThoughts struct {
	Username       string  `json:"username" binding:"required" gorm:"primaryKey"`
	MovieUUID      string  `json:"movie_uuid" binding:"required" gorm:"primaryKey"`
	PersonalRating float64 `json:"personal_rating" binding:"required" gorm:"not null"`
	Notes          string  `json:"notes"`
}

type UserMovieThoughtsInput struct {
	Username       string  `json:"username" binding:"required"`
	MovieUUID      string  `json:"movie_uuid" binding:"required"`
	PersonalRating float64 `json:"personal_rating"`
	Notes          string  `json:"notes"`
}
