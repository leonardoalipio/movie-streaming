package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/leonardoalipio/movie-streaming/server/controllers"
	"github.com/leonardoalipio/movie-streaming/server/middleware"
)

func SetupProtectedRoutes(router *gin.Engine) {
	router.Use(middleware.AuthMiddleware())

	router.GET("/movie/:imdb_id", controllers.GetMovie())
	router.POST("/addmovie", controllers.AddMovie())
}
