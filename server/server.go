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
		router.POST("/login", views.SignIn)
	}

	movies := router.Group("/movies", views.IdentifyUser)
	{
		movies.GET("/", views.GetAllMovies)
		movies.GET("/:uuid", views.GetMovieById)
		movies.POST("/", views.CreateMovie)
		movies.DELETE("/:uuid", views.DeleteMovieById)
		movies.PATCH("/:uuid", views.UpdateMovieById)
	}

	personalMovies := router.Group("/personal", views.IdentifyUser)
	{
		personalMovies.GET("/movies", views.GetMyMovies)
		//personalMovies.GET("/movies/:uuid", views.GetMyMovieById)
		//personalMovies.POST("/movies", views.AddMovieToCollection)
		//personalMovies.DELETE("/movies", views.DeleteMovieFromCollection)
		//personalMovies.PATCH("/movies/:uuid", views.UpdateMovieFromCollection)
	}

	_ = router.Run()
}
