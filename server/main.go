package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/leonardoalipio/movie-streaming/server/routes"
)

func main() {

	router := gin.Default()

	routes.SetupUnprotectedRoutes(router)
	routes.SetupProtectedRoutes(router)

	fmt.Println("Server is running...")
	router.Run(":8080")
}
