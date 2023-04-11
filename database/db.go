package database

import (
	"fmt"
	"github.com/duramash/go-movies/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "Ebalokratosa.228"
	dbname   = "postgres"
)

func GetDB() *gorm.DB {
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", user, password, host, port, dbname)
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		panic("no connection")
	}
	_ = db.AutoMigrate(&models.Movie{})
	return db
}
