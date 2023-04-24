package server

import (
	"github.com/duramash/go-movies/database"
	views "github.com/duramash/go-movies/views"
	"github.com/gin-gonic/gin"
)

var db = database.GetDB()

func RunServer() {
	router := gin.Default()

	router.Group("/auth")
	{
		router.POST("/register", views.SignUp)
		router.POST("/login", views.Login)
	}

	movies := router.Group("/movies", views.IdentifyUser)
	{
		movies.GET("/", views.GetAllMovies)
		movies.GET("/:uuid", views.GetMovieById)
		movies.POST("/", views.CreateMovie)
		movies.DELETE("/:uuid", views.DeleteMovieById)
		movies.PATCH("/:uuid", views.UpdateMovieById)
	}

	_ = router.Run()
}
