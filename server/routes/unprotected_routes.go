package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/leonardoalipio/movie-streaming/server/controllers"
)

func SetupUnprotectedRoutes(router *gin.Engine) {
	router.GET("/movies", controllers.GetMovies())
	router.POST("/register", controllers.RegisterUser())
	router.POST("/login", controllers.LoginUser())
}
