package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/leonardoalipio/movie-streaming/server/controllers"
)

func main() {

	router := gin.Default()

	router.GET("/movies", controllers.GetMovies())
	router.GET("/movie/:imdb_id", controllers.GetMovie())

	fmt.Println("Server is running...")
	router.Run(":8080")

}
