package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	fmt.Println("Server is running...")
	router.Run(":8080")

}
