package views

import (
	"fmt"
	"github.com/duramash/go-movies/database"
	"github.com/duramash/go-movies/models"
	"github.com/duramash/go-movies/util"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

var DB = database.GetDB()

func GetAllMovies(ctx *gin.Context) {
	var movies []models.Movie
	DB.Find(&movies)
	ctx.JSON(http.StatusOK, movies)
}

func CreateMovie(ctx *gin.Context) {
	movie := &models.Movie{}
	movie.UUID = uuid.New().String()
	if err := ctx.BindJSON(&movie); err != nil {
		util.HandleRequestError(ctx, http.StatusBadRequest, err)
		return
	}
	if movie.Rating < 0 || 10 < movie.Rating {
		util.HandleRequestError(ctx, http.StatusBadRequest, fmt.Errorf("rating is not in range between 0 and 10"))
		return
	}
	if err := DB.Create(movie).Error; err != nil {
		util.HandleRequestError(ctx, http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, movie)
}

func GetMovieById(ctx *gin.Context) {
	movie := models.Movie{}
	id := ctx.Param("uuid")
	if err := DB.First(&movie, "uuid = ?", id).Error; err != nil {
		util.HandleRequestError(ctx, http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusOK, movie)
}

func DeleteMovieById(ctx *gin.Context) {
	id := ctx.Param("uuid")
	movie := models.Movie{}
	if err := DB.First(&movie, "uuid = ?", id).Error; err != nil {
		util.HandleRequestError(ctx, http.StatusBadRequest, err)
		return
	}
	DB.Unscoped().Where("uuid = ?", id).Delete(&models.Movie{})
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("movie with uuid %s has been successfully deleted", id),
	})
}

func UpdateMovieById(ctx *gin.Context) {
	id := ctx.Param("uuid")
	movie := models.Movie{}
	updates := models.UpdateMovie{}
	if err := DB.First(&movie, "uuid = ?", id).Error; err != nil {
		util.HandleRequestError(ctx, http.StatusNotFound, err)
		return
	} // get movie from DB
	if err := ctx.BindJSON(&updates); err != nil {
		util.HandleRequestError(ctx, http.StatusBadRequest, err)
		return
	} // get request body
	if err := DB.Model(&movie).Updates(updates).Error; err != nil {
		util.HandleRequestError(ctx, http.StatusBadRequest, err)
		return
	} // apply updates from request body to movie from DB
	ctx.JSON(http.StatusOK, movie)
}
